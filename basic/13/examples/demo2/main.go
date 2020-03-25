package main

/*
#cgo CFLAGS: -I./clib
// 这里如果使用gcc编译动态链接库名不是lib开头的提示找不到 "ld: library not found for -ldemo"
#cgo LDFLAGS: -L./clib -l demo
#include "demo.h"
*/
import "C"// 切勿换行再写这个

import "fmt"

func main() {
	a := C.int(20)
	b := C.int(2)

	fmt.Println("go.call.c.file.demo")
	fmt.Printf("cAdd(%d, %d)=%d \n", a, b, C.Add(a, b))
}
/**
go run examples/demo2/main.go
go.call.c.file.demo
cAdd(20, 2)=22

 */