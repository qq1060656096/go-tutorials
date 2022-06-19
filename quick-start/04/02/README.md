# 常用同步原语

```
1. sync.Mutex 互斥锁
2. sync.RWMutex 读写锁
3. sync.WaitGroup 等待一组goroutine完成工作
4. sync.Once 无论调用多少次Do方法，都只执行一次。
5. atomic.Value 提供原子获取值和保存值
```

## 1. sync.Mutex 互斥锁

```
互斥锁：一段代码使用了互斥锁加锁，只会有一个goroutine执行它，其他执行这段代码的goroutine都会堵塞等待。

保证同一时刻内，只有一个 goroutine 持有锁，其他 goroutine 都要等待持有锁的 goroutine 释放锁后，才能继续执行。持有锁的 goroutine 没有释放，其他等待加锁 goroutine都会堵塞。在同一个 goroutine 内 二次加锁会导致死锁。



通过下面的2个例子我们验证：

结论1： goroutine 2 必须等待 goroutine 1 锁释放了，才能获取到锁，执行后续操作。

结论2: 在同一个 goroutine 内 二次加锁会导致死锁。
```

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁：保证同一时刻内，保证只有一个 goroutine 持有锁，其他 goroutine 都要等待持有锁的 goroutine 释放锁后，才能继续执行。
// 示例：goroutine 2 必须等待 goroutine 1 锁释放了，才能获取到锁，执行后续操作。
var mutex sync.Mutex

func main() {
	// goroutine 1
	fmt.Println("goroutine 1 before")
	mutex.Lock()
	fmt.Println("goroutine 1 locked")

	// goroutine 2
	go func() {
		fmt.Println("goroutine 2 lock before")
		// 这里会堵塞，一直到 goroutine 1 释放锁，goroutine 2才能继续执行
		mutex.Lock()
		fmt.Println("goroutine 2 locked")
		mutex.Unlock()
		fmt.Println("goroutine 2 unlock")
	}()
	for i := 1; i < 3; i++ {
		time.Sleep(time.Second * 1)
	}
	mutex.Unlock()
	fmt.Println("goroutine 1 unlock")
	time.Sleep(time.Second * 2)
}

/**
goroutine 1 before
goroutine 1 locked
goroutine 2 lock before
goroutine 1 unlock
goroutine 2 locked
goroutine 2 unlock
*/
```


```go
package main

import (
	"fmt"
	"sync"
)


var mutex sync.Mutex

func main() {
	// goroutine 1
	fmt.Println("goroutine 1 before")
	mutex.Lock()
	fmt.Println("goroutine 1 locked")
  // 在同一个 goroutine 内 二次加锁会导致死锁。
	mutex.Lock()
	mutex.Unlock()
	fmt.Println("goroutine 1 unlock")
}

/**
goroutine 1 before
goroutine 1 locked
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_SemacquireMutex(0x118c234, 0x0, 0x1)
        /usr/local/Cellar/go/1.16/libexec/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0x118c230)
        /usr/local/Cellar/go/1.16/libexec/src/sync/mutex.go:138 +0x105
sync.(*Mutex).Lock(...)
        /usr/local/Cellar/go/1.16/libexec/src/sync/mutex.go:81
main.main()
        /Users/concurrency/value.go:16 +0x1a5
 */
```

## 2. sync.RWMutex 读写锁
> 读写锁：读锁写锁互斥，有读锁时，写锁会堵塞，有写锁时读堵塞，只有读锁不会堵塞，常用于读多写少的场景

```go
Lock/Unlock 针对写操作：
  1. 不管有其他 goroutine 持有读锁还是 写锁，Lock 会一直堵塞，
  直到 Unlock/RUnlock 来释放锁
RLock/RUnlock  针对读操作：
  1. 当其他 goroutine持有读锁时，RLock直接返回
  2. 当有 goroutine 持有写锁时 RLock会一直堵塞，直到能获取锁
```


```go
package main

import (
	"sync"
	"time"
)

/*
读写锁：读锁写锁互斥，有读锁时，写锁会堵塞，有写锁时读堵塞，只有读锁不会堵塞，常用语读多写少的场景

Lock/Unlock 针对写操作：
  1. 不管有其他 goroutine 持有读锁还是 写锁，Lock 会一直堵塞，直到 Unlock/RUnlock 来释放锁

RLock/RUnlock  针对读操作：
  1. 当其他 goroutine持有读锁时，RLock直接返回
  2. 当有 goroutine 持有写锁时 RLock会一直堵塞，直到能获取锁
 */
var rwMutex sync.RWMutex

func main() {
	// goroutine 1
	fmt.Println("goroutine 1 before")
	rwMutex.RLock()
	fmt.Println("goroutine 1 locked")
	go func(){
		fmt.Println("goroutine 2 before")
		// goroutine 2 获取写锁会堵塞，只有 goroutine 1 释放锁以后， goroutine 2 才会继续执行
		rwMutex.Lock()
		fmt.Println("goroutine 2 locked")
		for i := 1; i < 3; i++ {
			time.Sleep(time.Second * 1)
		}
		rwMutex.Unlock()
		fmt.Println("goroutine 2 unlock")
	}()
	go func(){
		// 休眠1秒让 goroutine 2 获得写锁
		time.Sleep(time.Second * 1)
		fmt.Println("goroutine 3 before")
		// 这里会堵塞，等待 goroutine 2解锁后才能获得锁
		rwMutex.RLock()
		fmt.Println("goroutine 3 locked")
		rwMutex.RUnlock()
		fmt.Println("goroutine 3 unlock")
	}()
	for i := 1; i < 3; i++ {
		time.Sleep(time.Second * 1)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("goroutine 1 sleep 5 second")
	// goroutine 1 解锁后， goroutine 2才能获得写锁
	rwMutex.RUnlock()
	fmt.Println("goroutine 1 unlock")
	time.Sleep(time.Second * 10)

}

/**
goroutine 1 before
goroutine 1 locked
goroutine 2 before
goroutine 3 before
goroutine 1 sleep 5 second
goroutine 1 unlock
goroutine 2 locked
goroutine 2 unlock
goroutine 3 locked
goroutine 3 unlock

*/
```

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// 读锁之间互不影响
var rwMutex sync.RWMutex

func main() {
	// goroutine 1
	fmt.Println("goroutine 1 before")
	rwMutex.RLock()
	fmt.Println("goroutine 1 locked")
	go func(){
		fmt.Println("goroutine 2 before")
		rwMutex.RLock()
		fmt.Println("goroutine 2 locked")
		rwMutex.RUnlock()
		fmt.Println("goroutine 2 unlock")
	}()
	fmt.Println("goroutine 1 locking 1")
	time.Sleep(time.Second * 1)
	fmt.Println("goroutine 1 locking 2")

	rwMutex.RUnlock()
	fmt.Println("goroutine 1 unlock")
	time.Sleep(time.Second * 1)

}

/**
goroutine 1 before
goroutine 1 locked
goroutine 1 locking 1
goroutine 2 before
goroutine 2 locked
goroutine 2 unlock
goroutine 1 locking 2
goroutine 1 unlock

*/
```

## 3.sync.WaitGroup 等待一组goroutine完成工作

```
1. Add()设置等待多少个 goroutine
2. Done() goroutine完成工作调用 Done表示已经完成了
3. Wait()方法会一直堵塞，直到 所有的 goroutine 完成才会返回
```

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main()  {
	// 等待 4个 goroutine
	wg.Add(4)
	go func() {
		fmt.Printf("goroutine 0\n")
		for i := 1; i <= 3; i ++ {
			go func(index int) {
				fmt.Printf("goroutine %d\n", index)
				// goroutine完成时调用 Done
				wg.Done()
			}(i)
		}
		wg.Done()
	}()
	// Wait会一直堵塞，直到 所有的 goroutine 完成才会返回
	wg.Wait()
}
/**
goroutine 0
goroutine 3
goroutine 2
goroutine 1
 */

```

## 4. sync.Once 无论调用多少次Do方法，都只执行一次。

> 无论调用多少次Do方法，都只执行一次，并且所有 goroutine 会等待执行Do的goroutine完成后才会继续执行，否则会一直堵塞等待。适用于单例、懒加载等

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var once sync.Once

func main() {

	wg.Add(4)
	go func() {
		fmt.Printf("goroutine 0\n")
		for i := 1; i <= 3; i ++ {
			go func(index int) {
				fmt.Printf("goroutine %d\n", index)
				// 只会执行一次
				once.Do(func() {
					time.Sleep(time.Second * 2)
					fmt.Println("print once")
				})
				// 所有 goroutine 会等待执行Do的goroutine完成后才会继续执行，否则会一直等待
				fmt.Printf("goroutine %d after\n", index)

				// goroutine完成时调用 Done
				wg.Done()
			}(i)
		}
		wg.Done()
	}()
	// Wait会一直堵塞，直到 所有的 goroutine 完成才会返回
	wg.Wait()
}
/**
goroutine 0
goroutine 3
goroutine 1
goroutine 2
print once
goroutine 1 after
goroutine 3 after
goroutine 2 after
 */
```

## 5. atomic.Value 提供原子获取值和保存值

> atomic.Value 提供原子获取值和保存值并且是线程安全无锁实现的，比互斥锁和读写锁有更高的性能，适用于读多写少。适用于配置、热点数据、运营类数据加载。

```go
Store()：保存值
Load()：获取值
适用场景：
1. 配置加载
2. 热点数据
3. 运营类数据
我们在服务内部开启一个goroutine每隔10秒加载一次配置数据、热点数据、运营类数据等，
后续使用直接从内存里面拿就好。
```

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var hotDataOrConfig atomic.Value
var wg sync.WaitGroup
func main()  {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 20000;i++ {
			// 保存数据（示例：配置数据 热点数据 运营类数据）
			hotDataOrConfig.Store(i)
		}
	}()
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			go func(index int) {
				defer wg.Done()
				// 获取数据（示例：配置数据 热点数据 运营类数据）
				v := hotDataOrConfig.Load()
				fmt.Println("goroutine ", index, v)
			}(i)
		}
	}()
	wg.Wait()
}
```
