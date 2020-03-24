package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan string)
	// 通道发送者准备好了, 但是接受放没准备好, 所以会堵塞导致 所有goroutine 休眠产生死锁
	// 如果我们把"go cancelDeadlock()"这行代码注释去掉, 这里还有其他 goroutine 就不会产生死锁,
	// go cancelDeadlock()
	ch <- "no buffer channel"
	s := <- ch
	fmt.Println(s)
}

func cancelDeadlock() {
	for i := 0; i < 10000; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("cancelDeadlock.i=", i)
	}
}
/*
go run examples/demo3/main.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /go-tutorials/basic/11/examples/demo3/main.go:9 +0x59
exit status 2
*/