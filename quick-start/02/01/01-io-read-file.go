package main

import (
	"fmt"
	"io"
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

	// 读取文件所有内容
	fc, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	println("file content: ", string(fc))
}

