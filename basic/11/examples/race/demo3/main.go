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