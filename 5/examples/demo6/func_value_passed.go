package main

import "fmt"

func main()  {
	var i int
	var pi *int

	pi = &i
	fmt.Printf("%16s pi=%d \n", "", &pi)
	pointValuePassed(pi)

	fmt.Println()
	fmt.Printf("%17s i=%d \n", "", &i)
	normalValuePassed(i)
}

func pointValuePassed(pi *int)  {
	fmt.Printf("pointValuePassed.pi=%d \n", &pi)
}

func normalValuePassed(i int)  {
	fmt.Printf("normalValuePassed.i=%d \n", &i)
}

/*
pointValuePassed外的pi和pointValuePassed内的pi指针地址不一样证明是值传递
                 pi=824634335240
pointValuePassed.pi=824633778200

normalValuePassed外的i和normalValuePassed内的i指针地址不一样证明是值传递
                  i=824634417152
normalValuePassed.i=824633827512
*/