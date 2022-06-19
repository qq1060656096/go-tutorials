package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	_, fileName, _, _ := runtime.Caller(0)
	fmt.Println("当前路径：", fileName)
	// 打开文件
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// 读取 12 个字节
	b := make([]byte, 13)
	f.Read(b)
	fmt.Println(string(b))
}


