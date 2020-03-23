## 12. 常见的错误与陷阱

```
1. 误用短声明导致变量被覆盖
2. 误用字符串
3. 误用defer
4. 误用map
5. 何时使用new()和make()函数
6. 误用new()函数
7. 误用指针
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
> map是一种HashMap是无序的
> 多次循环迭代顺序是不固定的：不同的实现方法会使用不同的散列算法，得到不同的元素顺序。所以我们认为
> 这种顺序是随机的，go语言map估计这样设计的，这样可以使得程序在不同的散列算法实现下变得健壮。
> 如果需要按照某种排序来遍历map中的元素，我们必须显示的给键排序。

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

###### 7. 误用指针
> 请不要使用空指针传递参数容易引起恐慌panic

```go
// go run examples/demo7/pointer.go
package main

import "fmt"

func main()  {
	// 1.main.pi指针未空指针，
	// 2.传递到testPointer中时testPointer.pi也是空指针
	// 3.在testPointer方法改变了testPointer.pi指针的指向但是并没有改变main.pi指针的指向,
	// 4. main.pi指针仍然是空指针, 因为go语言是值传递
	// pi空指针
	var pi *int
	fmt.Printf("%9s pi=&%d \n", "", &pi)
	testPointer(pi)
	// 使用空指针会引起panic
	fmt.Printf("testPoint.after.pi=%d \n", *pi)
}

func testPointer(pi *int) {
	fmt.Printf("testPoint.pi=&%d \n", &pi)
	var i int = 10
	pi = &i
	fmt.Printf("testPoint.pi=%d \n", *pi)
}
/*
$ go run examples/demo7/pointer.go
          pi=&824634302480
testPoint.pi=&824634302496
testPoint.pi=10
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x109ac28]

goroutine 1 [running]:
main.main()

*/
```