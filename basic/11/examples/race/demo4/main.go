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