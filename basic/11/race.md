## 竞态示例

```
1. 示例:互斥锁
2. 示例:读写锁
3. 示例:死锁 
  3.1 示例:乱序加锁死锁
  3.2 示例:2次加锁死锁
  3.3 示例:无缓冲通道死锁
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