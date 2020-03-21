package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		fmt.Printf("1.key:%s, value:%d\n", k, v)
	}
	fmt.Println()
	for k, v := range m {
		fmt.Printf("2.key:%s, value:%d\n", k, v)
	}
}

/*
go run 13/examples/demo4/map.go
1.key:c, value:3
1.key:a, value:1
1.key:b, value:2

2.key:a, value:1
2.key:b, value:2
2.key:c, value:3
*/
