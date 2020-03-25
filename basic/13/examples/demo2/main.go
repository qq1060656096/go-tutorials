package main

/*
#cgo CFLAGS: -I lib
#cgo LDFLAGS: -L lib -l demolib
#include "demo_lib.h"
 */
import "C"// 切勿换行再写这个

import "fmt"

func main() {
	a := C.int(1)
	b := C.int(2)

	fmt.Println("go.call.c.file.demo")
	fmt.Printf("cAdd(%d, %d)=%d \n", a, b, C.Add(a, b))
}