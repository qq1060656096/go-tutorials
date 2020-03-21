// go run examples/demo3/defer.go
package main

import (
	"fmt"
)

func main() {
	testDefer()
}

func testDefer()  {
	fmt.Println("for start")
	for i := 1; i < 5; i ++ {
		// defer 先进后出，并且函数仅在函数返回时才会执行
		defer fmt.Println(i)
		fmt.Println("for doing", i)
	}
	fmt.Println("for end")
}