## 12. 常见的错误与陷阱

```
1. 误用短声明导致变量被覆盖
2. 误用字符串
3. 误用defer
4. 误用map
5. 何时使用new()和make()函数
6. 误用new()函数
```

###### 1. 误用短声明导致变量被覆盖
```
// go run examples/demo1/short_var.go
// if中语句中声明的name,age覆盖了if外声明的name,age变量
package main

import "fmt"

func main()  {
	var name , age = "张三", 18
	if name, age := "李四", 10; age > 10 {
		fmt.Println("if name=%s, age=%d", name, age)
	} else {
		fmt.Println("else name=%s, age=%d", name, age)// else name=%s, age=%d 李四 10
	}
	fmt.Println("name=%s, age=%d", name, age)// name=%s, age=%d 张三 18
}
```

###### 2.误用字符串
> 当对一个字符串进行频繁的操作时，请记住在go语言中字符串是不可变得。
> 使用+拼接字符串会导致拼接后的新字符串和之前字符串不同，导致需要分配新的存储空间存放新字符串
> 从而导致大量的内存分配和拷贝。频繁操作字符串建议用bytes.Buffer
````
// go run examples/demo2/string.go
package main

import (
	"fmt"
)

func main()  {
	var s = "test"
	for i := 0; i < 10; i++{
		// 字符串不可变，由于s字符串和si字符串拼接后字符串不在相同，
		// 导致需要分配新的存储空间存放新字符串，从而导致大量的内存分配和拷贝。
		si := fmt.Sprintf(" %d", i)
		s = s + si
	}
	fmt.Println(s)
}
````

###### 3.误用defer

```go
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
```
**结果**
```
go run examples/demo3/defer.go
for start
for doing 1
for doing 2
for doing 3
for doing 4
for end
4
3
2
1
```

###### 4.误用map
> map是一种HashMap是无序的，并且多次循环迭代顺序不一致(go底层实现了随机桶位置)

```go
// go run 13/examples/demo4/map.go
package main

import "fmt"

func main()  {
	m := map[string]int {
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v:= range m {
		fmt.Printf("1.key:%s, value:%d\n", k, v)
	}
	fmt.Println()
	for k, v:= range m {
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
```

###### 5. 何时使用new()和make()函数
```go
new: 只分配内存
make: 分配内存 和 初始化只能用于 slice、map、chan
切片(slice)、映射(map)、通道(chan)使用make
数组(array)、结构体(struct)、值类型使用new
```

###### 6.误用new()函数
```go
// go run 13/examples/demo6/new.go
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
```