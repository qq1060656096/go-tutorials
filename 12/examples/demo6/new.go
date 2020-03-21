package main

import (
	"fmt"
)

func main() {
	// new(T)函数返回T类型的指针并返回零值，切片本身就是指针类型并且切片的零值就是nil
	// new([]int) 相当于声明切片(即: var slice1 []int)
	slice1 := new([]int)
	fmt.Printf("slice1=%v\n", slice1)
	fmt.Printf("slice1==nil:%v\n", slice1 == nil)
}

/*
slice1=&[]
slice1==nil:false
*/
