package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Print("Go Goroutine Hello World!")

	// 主 goroutine(协程)休眠1秒, 让 go 语句创建的 goroutine(协程)有足够的时间执行
	// 所以正常输出
	time.Sleep(1 * time.Second)
}
/**
输出以下内容:
Go Goroutine Hello World!%
 */