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