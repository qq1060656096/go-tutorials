package main

import "fmt"

func main()  {
	// 1.main.pi指针未空指针，
	// 2.传递到testPointer中时testPointer.pi也是空指针
	// 3.在testPointer方法改变了testPointer.pi指针的指向但是并没有改变main.pi指针的指向,
	// 4. main.pi指针仍然是空指针, 因为go语言是值传递
	// pi空指针
	var pi *int
	fmt.Printf("%9s pi=&%d \n", "", &pi)
	testPointer(pi)
	// 使用空指针会引起panic
	fmt.Printf("testPoint.after.pi=%d \n", *pi)
}

func testPointer(pi *int) {
	fmt.Printf("testPoint.pi=&%d \n", &pi)
	var i int = 10
	pi = &i
	fmt.Printf("testPoint.pi=%d \n", *pi)
}
/*
$ go run examples/demo7/pointer.go
          pi=&824634302480
testPoint.pi=&824634302496
testPoint.pi=10
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x109ac28]

goroutine 1 [running]:
main.main()

*/