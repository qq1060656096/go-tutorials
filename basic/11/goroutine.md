## goroutine(协程)

> go语言在操作系统提供的内核线程只上, 搭建了一个特有的两级线程模型


```
1. 什么是goroutine(协程)
2. goroutine(协程) Hello World示例
3. goroutine常用函数 
```

### 1. 什么是goroutine(协程)
```
goroutine(协程) 是一种非常轻量级的实现, 可在单个进程里执行成千上万的并发任务.它是Go语言并发设计的核心.
goroutine(协程) 其实就是线程, 但是它比线程更小,十几个 goroutine(协程) 可能体现在底层就是五六个线程,
而且Go语言内部也实现了 goroutine(协程) 之间的内存共享. 

goroutine 是go 中最基本的执行单元.事实上每一个 go 程序至少有一个goroutine: 主 goroutine.
当程序启动时,它会自动创建. 
当主 goroutine退出后, 当前程序中所有的goroutine都会结束

在go里面, 每个并发执行的单元称为goroutine(协程).
当一个程序启动时,只有一个goroutine来调用main函数, 它称为主goroutine.新的goroutine通过go关键词进行创建.

语法上一个go语句时在普通函数或者方法调用前加上go关键字前缀. go语句本身会立即执行完成
go 函数名()
```

### 2. goroutine(协程) Hello World示例

> 以下这段代码什么也不输出, 这是为什么
```go
// examples/demo1/main.go
// go run examples/demo1/main.go 
package main

import "fmt"

func main()  {
	// 当主 goroutine退出后, 当前程序中所有的goroutine都会结束
	// go语句后面没有任何代码, 所以主 goroutine(协程) 总是比 go语句创建的 goroutine(协程) 更早的结束
	// 所以什么也不输出
	go fmt.Print("Go Goroutine Hello World!")
}
```
**更正后**

```go
// examples/demo2/main.go
// go run examples/demo2/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Print("Go Goroutine Hello World!")

	// 主 goroutine(协程)休眠1秒, 让 go 语句创建的 goroutine(协程)有足够的时间执行
	// 所以正常输出
	time.Sleep(1 * time.Second)
}
/**
输出以下内容:
Go Goroutine Hello World!%
 */
```

### 3. goroutine(协程)常用函数与结构体

```go
runtime.NumCPU()
runtime.GOMAXPROCS()
runtime.Goexit()
runtime.Gosched()
runtime.NumGoroutine()

runtime.LockOSThread()
runtime.UnlockOSThread()

// 配置
debug.SetMaxStack()
debug.SetMaxThreads()
// 垃圾回收
runtime.GC()
debug.FreeOSMemory()

sync.WaitGroup{}// WaitGroup等待goroutine的集合完成
sync.Mutex{}// 互斥锁
sync.RWMutex{}// 读写锁
```
