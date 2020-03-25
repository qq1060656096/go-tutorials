package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	fmt.Println("main.goroutine.channel.send.before")
	ch <- "send message"
	fmt.Println("receive: ", <- ch)
	fmt.Println("main.goroutine.channel.receive.after")
}
/**
go run examples/race/demo5/main.go
main.goroutine.channel.send.before
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        examples/race/demo5/main.go:10 +0xc4
exit status 2

 */