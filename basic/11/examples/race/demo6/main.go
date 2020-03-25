package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	Length int
	MaxLength int
	Data []string

}
func main() {
	goroutineCount := 2
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	hasResource := new(bool)
	wg := sync.WaitGroup{}
	wg .Add(goroutineCount + 1)
	for i := 0; i < goroutineCount; i ++ {
		go func(i int) {
			for j := 0; ; j ++ {
				fmt.Printf("goroutine.%d.doing.%d.start \n", i, j)
				mutex.Lock()
				if *hasResource {
					fmt.Printf("goroutine.%d.doing.%d.resource \n", i, j)
					*hasResource = false
				} else {
					fmt.Printf("goroutine.%d.doing.%d.resource.wait \n", i, j)
					cond.Wait()
				}
				mutex.Unlock()
				fmt.Printf("goroutine.%d.doing.%d.end \n", i, j)
				time.Sleep(2 * time.Second)
			}
			wg.Done()
		}(i)
	}

	go func() {
		for i := 0; ;i++ {
			mutex.Lock()
			*hasResource = true
			fmt.Printf("change.%d.hasResource \n", i)
			mutex.Unlock()
			cond.Signal()
			time.Sleep(2 * time.Second)
		}
	}()
	wg.Wait()
}

