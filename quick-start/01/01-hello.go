// 定义包名未 main 包，注意 main 是 Go 语言的一个特殊的包
package main

// 导入包 "fmt" 包，目的告诉 Go 编译器程序用到 了"fmt" 包
import "fmt"

// 程序执行的的开始函数，main()函数是可执行程序的入口函数
func main() {
	// 打印字符到控制台
	fmt.Println("say hello")
}
