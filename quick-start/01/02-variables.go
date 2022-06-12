package main

import "fmt"


// 第1种，指定变量类型，声明后不赋值，使用该类型的零值作为默认值
// var name   type
// var 变量名 变量类型
var name string //string=""
//  多重赋值
var id, age int // int=0, int=0

// 第2种，根据值类型自动判定变量类型
// var name   type     = expression
// var 变量名 变量类型 = 表达
var name2Str string = "name2Str"
var name2 = "name2"
//  多重赋值
var id2, age2 int = 2, 20 // int=1, int=10)

// 全局变量声明
var (
	id4 = 4 // int=4
	name4 string // string=
	age4 int = 40
)

/**
对于数字值是0
对于布尔值是false
对于字符串是""(即空字符串)
对于接口和引用类型(slice, 指针, map, 通道, 函数)值是nil
 */

func main()  {
	// 第3种，":="短变量声明(注意短变量声明最少声明一个新的变量)
	id3, age3 := 3, 30
	name3 := "name3"

	fmt.Println("name: ", name)
	fmt.Println("id: ", id)
	fmt.Println("age: ", age)

	fmt.Println("")
	fmt.Println("id2: ", id2)
	fmt.Println("name2: ", name2)
	fmt.Println("name2Str: ", name2Str)
	fmt.Println("age2: ", age2)

	fmt.Println("")
	fmt.Println("id3: ", id3)
	fmt.Println("name3: ", name3)
	fmt.Println("age3: ", age3)

	fmt.Println("")
	fmt.Println("id4: ", id4)
	fmt.Println("name4: ", name4)
	fmt.Println("age4: ", age4)
}

