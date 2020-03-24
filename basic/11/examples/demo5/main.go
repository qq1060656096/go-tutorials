package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan string, 1)
	wg.Add(2)// 等待2个 goroutine

	go send(ch)
	go receive(ch)

	wg.Wait() // 等待其他 goroutine 完成
}

func send(r chan string)  {
	fmt.Println("send.message.start")
	r <- "send message 1"
	r <- "send message 2"
	r <- "send message 3"
	close(r)
	fmt.Println("send.chanel.close")
	wg.Done()
}

func receive(c chan string) {
	for i := 0; ; i++ {
		s, ok := <- c

		if !ok {// false 通道关闭
			// 通道关闭, s字符是空字符
			fmt.Printf("receive.doing.%d=%s channel is close \n", i, s)
			wg.Done()
		} else {
			fmt.Printf("receive.doing.%d=%s \n", i, s)
		}
		time.Sleep(1 * time.Second)
	}
}
/*
go run examples/demo5/main.go
send.message.start
receive.doing.0=send message 1
send.chanel.close
receive.doing.1=send message 2
receive.doing.2=send message 3
receive.doing.3= channel is close
*/