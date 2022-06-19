# 如何优雅的使用 Goroutine：Goroutine 泄漏和排查

Go从语言层面就简化了并发编程，从而降低并发编程的难度，但这并不意味着我们可以随意使用Go关键字创建Goroutine。


之前接手一个下载服务，一个http请求内部开了好几个Goroutine还没有超时处理，大量的Groutine泄漏代码，各种秀代码，代码时写给人看的，越简单越好。接下来我们一起看看怎么优雅的使用 Goroutine



```go
1. 几种Goroutine常见的泄漏
  1.1.channel读写操作堵塞，由于逻辑问题导致一直堵塞
  1.2.Mutex/RWMutex 互斥锁、读写锁，由于逻辑问题导致一直堵塞
  1.3.goroutine 内的逻辑死循环
  1.4.goroutine 内的逻辑进入超长时间等待（比如调用第三方接口、库、cgo等没有设置超时控制）
2. 怎么避免Goroutine泄漏
3. 怎么排查Goroutine泄漏
4. 管控Goroutine生命周期
```

## 1. 几种Goroutine常见的泄漏

```
1.1.channel读写操作堵塞，由于逻辑问题导致一直堵塞
1.2.Mutex/RWMutex 互斥锁、读写锁，由于逻辑问题导致一直堵塞
1.3.goroutine 内的逻辑死循环
1.4.goroutine 内的逻辑进入超长时间等待（比如调用第三方接口、库、cgo等没有设置超时控制）
```

### 1. channel读写操作堵塞，由于逻辑问题导致一直堵塞
示例1：
```go
package main

import (
"fmt"
"runtime"
"time"
)

func search(name string) chan<- string {
	// ch是无缓冲通道，导致这个 goroutine 没人接收导致一直堵塞等待
	ch := make(chan string)
	go func() {
		fmt.Println(name, "send before")
		ch <- name // 一直会堵塞在这里，导致后面代码无法执行
		fmt.Println(name, "send done")
	}()
	return ch
}

func main() {
	for i := 1; i <= 3; i++ {
		go func(index int) {
			name := fmt.Sprintf("gourtine %d", index)
			search(name)
		}(i)
		fmt.Printf("goroutine count: %d\n", runtime.NumGoroutine())
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("end goroutine count: %d\n", runtime.NumGoroutine())
}
/**输出：
goroutine count: 2
goroutine count: 3
goroutine count: 4
gourtine 1 send before
gourtine 3 send before
gourtine 2 send before
end goroutine count: 4
 */
```

程序结束的时候还有4个goroutine，有3个还在堵塞等待有人从通道中接收数据，导致Goroutine泄漏。有人说很简单变成有缓冲通道就好，我把第11行代码加个1就好了。

```go
ch := make(chan string, 1)
```
修改以后执行结果如下：

```go
goroutine count: 2
goroutine count: 3
goroutine count: 4
gourtine 3 send before
gourtine 3 send done
gourtine 1 send before
gourtine 1 send done
gourtine 2 send before
gourtine 2 send done
end goroutine count: 1
```

虽然解决了问题，还是仍然不够好。

我们又修改了版本2代码如下：

```go
package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// context 应该 search 函数外面传入，这里为了演示方便，我们后面说怎么优雅的使用
func search(name string) chan<- string {
	ch := make(chan string)
	go func() {
		// 这个 goroutine 加上了超时控制
		// 要么有接收通道数据 或者 100毫秒后超时返回
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond * 100)
		defer cancel()
		fmt.Println(name, "send before")
		select {
		case ch <- name:
			fmt.Println(name, "send done")
		case <-ctx.Done():
			fmt.Println(name, "send timeout")
		}
	}()
	return ch
}

func main() {
	for i := 1; i <= 3; i++ {
		go func(index int) {
			name := fmt.Sprintf("gourtine %d", index)
			search(name)
		}(i)
		fmt.Printf("goroutine count: %d\n", runtime.NumGoroutine())
	}
	time.Sleep(time.Second * 5)
	fmt.Printf("end goroutine count: %d\n", runtime.NumGoroutine())
}
/**
goroutine count: 2
goroutine count: 3
goroutine count: 4
gourtine 3 send before
gourtine 2 send before
gourtine 1 send before
gourtine 3 send timeout
gourtine 2 send timeout
gourtine 1 send timeout
end goroutine count: 1
*/
```

程序结束的时候只有1个goroutine，因为3个工作的Goroutine等了100毫秒，没人接收，就超时返回了。我们一般在一个高频的http、rpc 请求都会做超时处理，要么降级处理返回热点数据，要么报错。

## 2. Mutex/RWMutex 互斥锁、读写锁，由于逻辑问题导致一直堵塞
示例2.1：Goroutine加锁不解锁就退出了，导致后续goroutine一直堵塞

```go
package main

import(
	"fmt"
	"runtime"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {
	for i := 1; i <= 3; i ++ {
		go func(index int) {
			name := fmt.Sprintf("goroutine %d", index)
			defer fmt.Println(name, " defer")
			fmt.Println(name, " lock before")
			// 第一个获取到锁的goroutine没有解锁就退出了，导致后续goroutine一直堵塞在这里
			mutex.Lock()
			fmt.Println(name, " locked")
		}(i)
		fmt.Println("goroutine count: ", runtime.NumGoroutine())
	}

	time.Sleep(time.Second * 2)
	fmt.Println("goroutine count: ", runtime.NumGoroutine())
}
/**
goroutine count:  2
goroutine count:  3
goroutine count:  4
goroutine 2  lock before
goroutine 2  locked
goroutine 2  defer
goroutine 1  lock before
goroutine 3  lock before
goroutine count:  3
*/

```

goroutine2 抢到了锁，但是它退出的时候没有解锁，导致goroutine1和goroutine3到程序结束的时候还在获取锁，一直堵塞着，导致goroutine泄露。我们在加锁后面增加“defer mutex.Unlock()”，后运行发现没有出现泄露

```go
goroutine count:  2
goroutine count:  3
goroutine count:  4
goroutine 1  lock before
goroutine 1  locked
goroutine 1  defer
goroutine 2  lock before
goroutine 2  locked
goroutine 2  defer
goroutine 3  lock before
goroutine 3  locked
goroutine 3  defer
goroutine count:  1
```

## 示例2.2：Goroutine加读锁没有解锁，就去加写锁。导致后续 Goroutine 在加读锁的时候会堵塞

```go
package main

import (
	"runtime"
	"sync"
	"fmt"
	"time"
)
var rwMutex sync.RWMutex

func main() {
	for i := 1; i <= 3; i ++ {
		go func(index int) {
			name := fmt.Sprintf("goroutine %d", index)
			defer fmt.Println(name, " defer")
			fmt.Println(name, " RLock before")
			// 第一个拿到读锁goroutine 没有解锁，就去加写锁。导致后续 goroutine 在加 读锁的时候会堵塞
			rwMutex.RLock()
			fmt.Println(name, " RLock locked")
			rwMutex.Lock()
			fmt.Println(name, " locked")
		}(i)
		fmt.Println("goroutine count: ", runtime.NumGoroutine())
	}

	time.Sleep(time.Second * 2)
	fmt.Println("goroutine count: ", runtime.NumGoroutine())
}
/**
goroutine count:  2
goroutine count:  3
goroutine count:  4
goroutine 3  RLock before
goroutine 3  RLock locked
goroutine 1  RLock before
goroutine 2  RLock before
goroutine count:  4

 */

```

goroutine3先拿到读锁，没有解锁就去加写锁，导致goroutine1和goroutine2在加读锁的时候一直获取不到堵塞，导致goroutine泄露。


### 3.goroutine 内的逻辑死循环

示例3：
```go
package main

import(
	"fmt"
	"runtime"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {
	for i := 1; i <= 3; i ++ {
		go func(index int) {
			name := fmt.Sprintf("goroutine %d", index)
			for {
				fmt.Println(name, " running")
				time.Sleep(time.Millisecond * 900)
			}
		}(i)
		fmt.Println("goroutine count: ", runtime.NumGoroutine())
	}

	time.Sleep(time.Second * 2)
	fmt.Println("goroutine count: ", runtime.NumGoroutine())
}
/**
goroutine count:  2
goroutine count:  3
goroutine count:  4
goroutine 3  running
goroutine 1  running
goroutine 2  running
goroutine 2  running
goroutine 1  running
goroutine 3  running
goroutine 3  running
goroutine 1  running
goroutine 2  running
goroutine count:  4
*/
```


##  4. goroutine 内的逻辑进入超长时间等待（比如调用第三方接口、库、cgo等没有设置超时控制）
示例4：

```go
package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func search(name string) {
	// 用不存在域名，模拟请求第三方接口超长时间等待
	res, err := http.Get("http://notfound.notfound.com")
	if  err != nil {
		fmt.Println(name, " err:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println(name, " ok")
}

func main() {
	for i := 1; i <= 3; i++ {
		go func(index int) {
			name := fmt.Sprintf("gourtine %d", index)
			search(name)
		}(i)
		fmt.Printf("goroutine count: %d\n", runtime.NumGoroutine())
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("end goroutine count: %d\n", runtime.NumGoroutine())
}
/**输出：
goroutine count: 2
goroutine count: 3
goroutine count: 4
end goroutine count: 10

*/

```
在调用第三方接口的时候，由于接口很慢，很久不返回结果。http.client默认是没有设置超时时间，导致Goroutine泄漏。所以在调用第三方接口的时候我们要做好超时控制。

### 2. 怎么避免Goroutine泄漏

```go
1.避免channel堵塞（在写channel入方关闭，不关闭也没事，gc会自动回收）
2.避免死锁堵塞（同类型的锁成对出现，固定顺序锁）
3.避免Goroutine内部逻辑死循环
4.不知道一个资源或者函数多久返回时，要做好超时控制
5.本质就是管控Goroutine生命周期
```
## 3. 怎么排查Goroutine泄漏
一般我们使用 go tool & Graph & 火焰图 & pprof 来排查 Goroutine泄漏，pprof结合业务代码确定是否泄漏。可以通过火焰图，看的那些函数占用的比较多，来排查。具体的使用请方式自行百度。



## 4. 管控Goroutine生命周期

```go
只有把Groutine的生命周期管理起来，才能避免Goroutine泄漏。那怎么管控Goroutine的生命周期，我们一起来看这篇文章如何管控Goroutine的生命周期
```
