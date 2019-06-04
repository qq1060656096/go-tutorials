package demo3

import "fmt"

// 成功的示例
func ExampleOk() {
	b := Login("root", "123456")
	fmt.Println(b)
	//Output:
	// true
}

// 失败的示例
func ExampleFail() {
	b := Login("root", "123456")
	fmt.Println(b)
	//Output:
	// false
}

// 成功的示例
func ExampleListOk() {
	b := Login("root", "123456")
	fmt.Println(b)
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	//Unordered Output:
	// true
	// 3
	// 2
	// 1
}