package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan string, 2)
	wg.Add(2)
	go send(ch)
	go receive(ch)
	wg.Wait()
}

func send(s chan <- string) {
	for i := 0; i < 5; i ++ {
		str := fmt.Sprintf("send.message.%d", i)
		fmt.Println("sender", i, ": ", str)
		s <- str
	}
	close(s)
	wg.Done()
}


func receive(r <- chan string) {
	for i := 0; i < 5; i++ {
		str := fmt.Sprintf("receiver.%d: %s", i, <-r)
		fmt.Println(str)
	}
	wg.Done()
}
/**
go run examples/demo6/main.go
sender 0 :  send.message.0
sender 1 :  send.message.1
sender 2 :  send.message.2
sender 3 :  send.message.3
receiver.0: send.message.0
receiver.1: send.message.1
receiver.2: send.message.2
receiver.3: send.message.3
sender 4 :  send.message.4
receiver.4: send.message.4
 */