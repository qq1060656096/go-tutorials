## 竞态示例

```
1. 示例:互斥锁
2. 示例:读写锁
3. 示例:死锁 
  3.1 示例:乱序加锁死锁
  3.2 示例:2次加锁死锁
  3.3 示例:无缓冲通道死锁
4. 示例:条件变量
```

### 1. 示例:互斥锁
```go
// go run examples/race/demo1/main.go
package main

import (
	"fmt"
	"sync"
)

var mutex = sync.Mutex{}
var wg = sync.WaitGroup{}
func main() {
	count := new(int)
	lpCount := 3
	for i := 0; i < lpCount;i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("goroutine.%d.lock count(%d) \n", i, *count)
			mutex.Lock()
			*count ++
			fmt.Printf("goroutine.%d.locked count(%d) \n", i, *count)
			mutex.Unlock()
			fmt.Printf("goroutine.%d.unlock count(%d) \n", i, *count)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("count=", *count)
}
/**
go run examples/race/demo1/main.go
goroutine.1.lock count(0)
goroutine.1.locked count(1)
goroutine.1.unlock count(1)
goroutine.2.lock count(0)
goroutine.2.locked count(2) 
goroutine.2.unlock count(2)
goroutine.0.lock count(0)
goroutine.0.locked count(3)
goroutine.0.unlock count(3)
count= 3

 */
```

### 2. 示例:读写锁
```go
// go run examples/race/demo2/main.go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}
var rwMutex = sync.RWMutex{}

func main() {
	count := new(int)
	lpCount := 5
	// 读锁
	for i := 0; i < lpCount; i ++ {
		go func(i int) {
			wg.Add(1)
			rwMutex.RLock()
			fmt.Printf("goroutine.read.%d: %d\n", i , *count)
			time.Sleep(100 * time.Millisecond)
			rwMutex.RUnlock()
			wg.Done()
		}(i)
	}
	// 写锁
	for i := 0; i < lpCount; i ++ {
		go func(i int) {
			wg.Add(1)
			rwMutex.Lock()
			*count ++
			fmt.Printf("goroutine.write.%d: %d\n", i , *count)
			rwMutex.Unlock()
			runtime.Gosched()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
/*
go run examples/race/demo2/main.go
goroutine.read.0: 0
goroutine.write.4: 1
goroutine.read.1: 1
goroutine.read.3: 1
goroutine.read.4: 1
goroutine.read.2: 1
goroutine.write.0: 2
goroutine.write.1: 3
goroutine.write.2: 4
goroutine.write.3: 5

*/
```

### 3. 示例:死锁

##### 3.1 示例:乱序加锁死锁
> 一个协程goroutine1.0.mutex1.locked, 但是它需要加锁 mutex2, 但是 mutex2 被加锁了，所以需要等待
> 另一个协程 goroutine2.0.mutex2.locked 但是它需要加锁 mutex1, 但是 mutex1 被加锁了，所以需要等待
> 两个协程相互等待但是没有结束, 造成死锁
```go
// go run --race examples/race/demo3/main.go
package main

import (
	"fmt"
	"sync"
)

var mutex1 = sync.Mutex{}
var mutex2 = sync.RWMutex{}
var wg = sync.WaitGroup{}

func main() {
	count1 := new(int)
	count2 := new(int)
	lpCount := 1
	wg.Add(2 * lpCount)

	for i:=0; i < lpCount ; i++ {
		go func(i int) {
			for j := 0;; j ++{
				mutex1.Lock()
				*count1 ++
				fmt.Printf("goroutine1.%d.mutex1.locked\n", i)
				fmt.Printf("goroutine1.%d.mutex2.lock.wait\n", i)
				mutex2.Lock()
				mutex1.Unlock()
				*count2 ++
				fmt.Printf("goroutine1.%d.mutex2.locked %d j=%d\n", i, *count2, j)
				mutex2.Unlock()
			}
			wg.Done()
		}(i)
	}

	for i:=0; i < lpCount ; i++ {
		go func(i int) {
			for j := 0; ; j++{
				mutex2.Lock()
				*count2 ++
				fmt.Printf("goroutine2.%d.mutex2.locked %d j=%d\n", i, *count2, j)
				fmt.Printf("goroutine2.%d.mutex1.lock.wait\n", i)
				mutex1.Lock()
				mutex2.Unlock()
				*count1 ++
				fmt.Printf("goroutine2.%d.mutex1.locked\n", i)
				mutex1.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
/**
一个协程goroutine1.0.mutex1.locked, 但是它需要加锁 mutex2, 但是 mutex2 被加锁了，所以需要等待
另一个协程 goroutine2.0.mutex2.locked 但是它需要加锁 mutex1, 但是 mutex1 被加锁了，所以需要等待
两个协程相互等待但是没有结束, 造成死锁


go run --race examples/race/demo3/main.go
goroutine1.0.mutex1.locked
goroutine2.0.mutex2.locked 1 j=0
goroutine2.0.mutex1.lock.wait
goroutine1.0.mutex2.lock.wait

 */
```

##### 3.2 示例:2次加锁死锁

```go
// go run examples/race/demo4/main.go
package main

import (
	"fmt"
	"sync"
)

var mutex = sync.Mutex{}

func main() {
	mutex.Lock()
	fmt.Println("mutex.lock.1")
	mutex.Lock()
	fmt.Println("mutex.lock.2")
	mutex.Unlock()
}
/**
go run examples/race/demo4/main.go
mutex.lock.1
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_SemacquireMutex(0x1192984, 0x0, 0x1)
 */
```

##### 3.3 示例:无缓冲通道死锁
```go
// go run examples/race/demo5/main.go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	fmt.Println("main.goroutine.channel.send.before")
	ch <- "send message"
	fmt.Println("receive: ", <- ch)
	fmt.Println("main.goroutine.channel.receive.after")
}
/**
go run examples/race/demo5/main.go
main.goroutine.channel.send.before
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        examples/race/demo5/main.go:10 +0xc4
exit status 2

 */
```

### 4. 示例:条件变量
```go
// go run examples/race/demo6/main.go

```