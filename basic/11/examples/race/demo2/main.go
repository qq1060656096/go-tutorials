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