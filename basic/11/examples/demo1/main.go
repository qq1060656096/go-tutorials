package main

import "fmt"

func main()  {
	// 当主 goroutine退出后, 当前程序中所有的goroutine都会结束
	// go语句后面没有任何代码, 所以主 goroutine(协程) 总是比 go语句创建的 goroutine(协程) 更早的结束
	// 所以什么也不输出
	go fmt.Print("Go Goroutine Hello World!")
}