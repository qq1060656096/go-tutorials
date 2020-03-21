package main

import "fmt"

func main() {
	var name, age = "张三", 18
	if name, age := "李四", 10; age > 10 {
		fmt.Println("if name=%s, age=%d", name, age)
	} else {
		fmt.Println("else name=%s, age=%d", name, age) // else name=%s, age=%d 李四 10
	}
	fmt.Println("name=%s, age=%d", name, age) // name=%s, age=%d 张三 18
}
