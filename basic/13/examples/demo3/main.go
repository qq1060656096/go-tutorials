package main

//#include "./c/demo.c"
import "C"// 切勿换行再写这个


import (
	"fmt"
)

func main() {
	a := C.int(3)
	b := C.int(30)

	fmt.Println("go.call.c.file.demo")
	fmt.Printf("Add(%d, %d)=%d \n", a, b, C.Add(a, b))
}
/**
go run examples/demo3/main.go
go.call.c.file.demo
cAdd(1, 2)=3
 */