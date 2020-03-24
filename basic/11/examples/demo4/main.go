package main

import (
	"fmt"
	"sync"
)

// 使用WaitGroup来等待2个goroutine结束, 主 goroutine才结束
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan string)
	wg.Add(2) // 等待2个 goroutine
	go receive(ch)
	go send(ch)
	wg.Wait() // 等待 goroutine 结束
}

func send(sc chan string) {
	fmt.Println("send.doing")
	s := "send message one"
	sc <- s
	close(sc)
	sc <- "close channel send message"// 这里会引起"panic: send on closed channel"
	wg.Done() // 完成1个 goroutine
}

func receive(rc chan string) {
	fmt.Println("receive.doing")
	fmt.Printf("receive: %s \n", <-rc)
	wg.Done() // 完成1个 goroutine
}

/**
 go run examples/demo4/main.go
receive.doing
send.doing
panic: send on closed channel

goroutine 34 [running]:
main.send(0xc00009c000)
        /basic/11/examples/demo4/main.go:24 +0xce
created by main.main
        /basic/11/examples/demo4/main.go:15 +0xa1
exit status 2

 */