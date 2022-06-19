# 如何优雅的使用Goroutine下-如何管控Goroutine的生命周期


通过之前的这篇Goroutine怎么泄漏文章，我们知道Goroutine是怎么泄漏的，接下来我们一起看下怎么管控Goroutine的生命周期。

只有我们把Groutine的生命周期管理起来，才能避免Goroutine泄漏。那怎么管控Goroutine的生命周期，我们做到以下几点就可以了。


在启动goroutine的时候你要问题自己3个问题

1. 尽量把并发扔给调用者（因为只有调用者才知道一个 Goroutine 什么开始什么时候结束）
2. Goroutine 什么时候结束（让调用者知道 Goroutine 什么时候结束）
3. 怎么控制 Goroutine 结束（让调用者知道怎么控制 Goroutine 结束）

示例1：
开启了2个Goroutine，判断了错误还记录了日志，看起来习惯挺好的。但是有点问题，如果http.ListenAndServe报错，主Goroutine一直堵塞导致无法退出

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hello word")
	})
	// 访问地址：http://localhost/hello
	go func() {
		err := http.ListenAndServe(":80", serverMux)
		if err != nil {
			log.Println(err)
		}
	}()

	//  访问地址：http://localhost:8001/debug/pprof/
	go func() {
		err := http.ListenAndServe(":8001", http.DefaultServeMux)
		if err != nil {
			log.Println(err)
		}
	}()
	// 阻塞
	select {}
}
```

修改版本1：有些小伙伴很聪明，不就是无法退出嘛，把log.Println()改成log.Fatal()。如果还有Goroutine执行其他逻辑，你无法保证程序的完整性。不仅导致defer无法执行，还会导致程序无法平滑退出。
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hello word")
	})
	// 访问地址：http://localhost/hello
	go func() {
		defer func() {
			log.Println("http app")
		}()
		err := http.ListenAndServe(":80", serverMux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	//  访问地址：http://localhost:8001/debug/pprof/
	go func() {
		defer func() {
			log.Println("http debug")
		}()
		err := http.ListenAndServe(":8001", http.DefaultServeMux)
		if err != nil {
			log.Fatal(err)
		}
	}()
	// 阻塞
	select {}
}
```

修改版本2：

第1步：如果一个方法干了2多件事，我们一般会拆分成多个方法。这里我们拆分成 httpApp()和httpDebug()

第2步：我们要控制Goroutine的生命周期，把并发扔给调用者，Goroutine什么时候结束、怎么结束

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

var wg = sync.WaitGroup{}
var stop chan error
var closeChanShutdownServer chan int

func main() {
	stop = make(chan error, 2)
	closeChanShutdownServer = make(chan int, 2)
	// 并发扔给调用者
	wg.Add(2)
	go func() {
		wg.Done()
		stop <- httpApp(closeChanShutdownServer)
	}()
	go func() {
		wg.Done()
		stop <- httpDebug(closeChanShutdownServer)
	}()

	shutdownClose := false
	for i := 0; i < cap(stop); i++ {
		err := <-stop
		log.Println("main: ", err)
		if !shutdownClose {
			shutdownClose = true
			close(closeChanShutdownServer)
		}
	}

	wg.Wait()
	log.Println("main.goroutine done")
}

// 访问地址：http://localhost/hello
func httpApp(closeChanShutdown <-chan int) error {
	defer func() {
		log.Println("http app")
	}()

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hello word")
	})
	return httpAppRun(":80", serverMux, closeChanShutdown)
}

//  访问地址：http://localhost:8001/debug/pprof/
func httpDebug(closeChanShutdown <-chan int) error {
	defer func() {
		log.Println("http debug")
	}()
	return httpAppRun(":8001", http.DefaultServeMux, closeChanShutdown)
}

// 运行http
// 什么时候goroutine结束：关闭通道 closeChanShutdown 后结束
// 怎么控制goroutine结束：关闭通过 closeChanShutdown 控制goroutine结束
func httpAppRun(addr string, handler http.Handler, closeChanShutdown <-chan int) error {
	wg.Add(1)
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		log.Println(addr, "goroutine shutdown waiting")
		<-closeChanShutdown
		log.Println(addr, "goroutine shutdown start")
		s.Shutdown(context.Background())
		log.Println(addr, "goroutine shutdown done")
		wg.Done()
	}()
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
	return err
}
```

示例2：定义了一个 search 函数，使用 200 毫秒超时模拟 http、rpc请求

```go
package main

import (
	"fmt"
	"time"
)

// 定义了一个 search 函数，使用 200 毫秒超时模拟 http、rpc请求
func search(term string) (string, error) {
     time.Sleep(200 * time.Millisecond)
     return "some value", nil
}

// process 方法 调用 search 方法找到记录并打印
// 它有什么问题呢，它会一直堵塞在这里，对于一些应用延迟是不可接受的并且可能是调用多个rpc
func process(term string) error {
     record, err := search(term)
     if err != nil {
		 return err
	 }
     fmt.Println("Received:", record)
     return nil
}

func main() {
	process("test")
}

```

修改版本1：process 方法 调用 search 方法找到记录并打印，如果超过 100 毫秒还没拿到结果就返回一个错误，看起来习惯挺好的, 看起来也没什么问题

search 函数 200毫秒超时，process 函数 100 秒超时，ch 通道是无缓冲的导致，必须等待通道接受者接收到了值，这里没有接收者，导致发送方无限堵塞等待。因为接收者 100ms 超时以及关闭到，导致 goroutine 泄漏。

有的人又说 这里可以设置有缓冲通道，我设置一个 1 不就好了。其实你是不知道search函数什么时候退，所以我们要做好超时控制。

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"
)


// 定义了一个 search 函数，使用 200 毫秒超时模拟 http、rpc请求
func search(term string) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return "some value", nil
}


type result struct {
	record string
	err    error
}


// process 方法 调用 search 方法找到记录并打印，如果超过 100 毫秒还么拿到结果就返回一个错误
// 看起来习惯挺好的, 它有什么问题呢
// search 函数 200毫秒超时，process 函数 100 秒超时，ch 通道是无缓冲的导致，必须等待通道接受者接收到了值，
// 这里没有接收者，因为接收者 100ms 超时以及关闭到，导致 goroutine 泄漏
// 有的人又说 这里可以设置有缓冲通道，我设置一个 1 不就好了。
// 其实你是不知道search函数什么时候退，所以我们要做好超时控制
func process(term string) error {
	// 创建一个100毫秒取消的 context
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// 创建一个通道接受 search 的返回值
	ch := make(chan result)
	//ch := make(chan result, 1)

	// 开启一个 goroutine 搜索记录
	// 用 result 把结果包起来，发送到 通道中
	// 这里有个好的地方，把并发扔给调用者
	go func() {
		fmt.Println("call search before")
		record, err := search(term)
		fmt.Println("call search after")
		ch <- result{record, err}
		fmt.Println("call search before sen chan")
	}()

	// 这里会堵塞等待
	// 要嘛 100毫秒超时，返回一个错误
	// 要嘛 从通道接受到值
	select {
	case <-ctx.Done():
		return errors.New("search canceled")
	case result := <-ch:
		if result.err != nil {
			return result.err
		}
		fmt.Println("Received:", result.record)
		return nil
	}
}

func main() {
	process("test")
  
	// 这里测试为了简单，请不要模仿
	for {
		runtime.Gosched()
	}
}

```

也就是说当我们不知道一个 Goroutine 什么时候结束时，我们不应该启动它。

只要做好这3点（把并发扔给调用者、Goroutine什么时候结束、怎么控制Goroutine结束），无论超时退出或者用channel告诉Goroutine什么时候退出，go关键字你就用好。

