## 通道(channel)
> go 奉行通过通信来共享内存,而不是共享内存来通信.

```
1. 什么是通道(channel)
2. 无缓冲通道和缓冲通道
3. 关闭通道问题 
4. 示例
  4.1 示例:为什么说无缓冲通道要求发送协程和接收协程同时准备好
  4.2 示例:向已关闭通道发送数据将会触发 panic
  4.3 示例:从已关闭的通道接收数据时将不会发生阻塞, 并且接收到的是通道类型的零值
  4.4 示例:单向通道
```


### 1. 什么是通道(channel)
```
channel是 goroutine(协程) 之间的通信机制. 一个channel是一个通信机制
goroutine(协程)之间可以通过它发送消息和接收消息
```

**通道(channel)目的**

```
通道(channel)目的: 负责协程之间的通信, 从而避开所有由共享内存导致的陷阱. 
这种通过通道进行通信的方式保证了同步性.
数据在通道中进行传递, 保证在任何时间, 一个数据被设计为只有一个协程可以对其访问, 所以不会发生数据竞争.
```

**通道使用**
```go
ch := make(chan int)
ch <- x // 写入数据到通道中
x = <- chan // 从通道中读取数据
close(ch) // 关闭通道
```

### 2. 无缓冲通道和缓冲通道
```
无缓冲通道: 通道容量为0的通道就是无缓冲通道
特点: 要求发送协程和接收协程同时准备好, 能完成发送和接收操作. 如果两个协程没有同时准备好, 通道会导致先执行发送或接收操作的协程阻塞等待

有缓冲通道: 通道容量大于0的通道就是缓冲通道, 它有类似于消息队列先进先出(FIFO), 队列的长度在创建的时候通过make 的容量参数来设置.
特点: 只有在通道是空的情况下, 接收操作才会阻塞. 只有在通道缓冲区容量满的情况下, 发送操作才会阻塞。

最简单方式调用make函数创建的是一个无缓存的channel, 但是我们也可以指定第二个整型参数,对应channel的容量.
如果channel的容量大于零，那么该channel就是带缓存的channel
```

**创建无缓冲通道**
```go
// 创建无缓冲通道
ch := make(chan int)
ch := make(chan int, 0)
```

**创建缓冲通道**
```go
ch := make(chan int, 1)
ch := make(chan int, 2)
```

### 3. 关闭通道问题 
```
1. 试图向已关闭通道发送数据将会触发 panic
2. 无论怎么样只有发送者才需要关闭通道, 接收者永远不需要关闭通道, 因为接收者通常无法判断发送者是否还会向该通道发送元素值.
3. 从已关闭的通道接收数据时将不会发生阻塞, 并且接收到的是通道类型的零值
4. 从通道取出数据时应该检测通道是否关闭
value, ok := ch
第二个结果ok用于检测通道是否关闭
```

### 4. 示例

##### 4.1 示例:为什么说无缓冲通道要求发送协程和接收协程同时准备好
```go
// go run examples/demo3/main.go
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
```

##### 4.2 示例:向已关闭通道发送数据将会触发 panic

```go
// go run examples/demo4/main.go
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
```
**示例: 双向无缓冲通道**
> 使用WaitGroup来等待2个goroutine结束, 主 goroutine才结束
```go
// go run examples/demo3/main.go
package main

import (
	"fmt"
	"sync"
)
// 使用WaitGroup来等待2个goroutine结束, 主 goroutine才结束
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan string)
	wg.Add(2)// 等待2个 goroutine
	go receive(ch)
	go send(ch)
	wg.Wait()// 等待 goroutine 结束
}

func send(sc chan string) {
	s := "test"
	sc<- s
	fmt.Printf("send: %s \n", s)
	wg.Done()// 完成1个 goroutine
}


func receive(rc chan string) {
	fmt.Printf("receive: %s \n", <-rc)
	wg.Done()// 完成1个 goroutine
}
```

##### 4.3 示例:从已关闭的通道接收数据时将不会发生阻塞, 并且接收到的是通道类型的零值

```go
// go run examples/demo5/main.go
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
```

##### 4.4 示例:单向通道

```go
// go run examples/demo6/main.go
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
```
