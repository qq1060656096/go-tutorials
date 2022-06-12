package main

import "fmt"

func main() {
	//1. 基本数据类型
	//    1.1 整数
	//    1.2 浮点
	//    1.3 复数
	//    1.4 布尔(boolean)
	//    1.5 字符串(string)
	// 	  1.6 字符(byte)
	var idUint64 uint64 = 1
	fmt.Println("基本数据类型：")
	fmt.Println("uint64：", idUint64)

	var f32 float32 = 2.0
	fmt.Printf("float32：%#v\n", f32)
	var complex641 complex64 = 1 + 2i
	fmt.Printf("complex64：%#v\n", complex641)
	var yes bool = true
	fmt.Printf("bool：%#v\n", yes)

	var s string = "hello 你好"
	//  字符串转字节序列
	b := []byte(s)
	fmt.Printf("byte：%#v\n", b)

	// 字符串转unicode码点
	runes := []rune(s)// []int32{'h', 'e', 'l',' l', 'o', ' ', '你', '好'}
	fmt.Printf("rune：%#v\n", runes)



	//2. 复合数据类型
	//    2.1 数组
	//    2.2 切片(slice)
	//    2.3 map
	//    2.4 结构体
	//3. 任意数据类型(空接口)
	//    interface{}
	//4. 类型断言
	//    type.(类型)
	//5. 类型转换
	//    type()
}
