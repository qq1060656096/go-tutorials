package main

import (
	"fmt"
)

func main()  {
	var s = "test"
	for i := 0; i < 10; i++{
		s = s + fmt.Sprintf(" %d", i)
	}
	fmt.Println(s)
}

