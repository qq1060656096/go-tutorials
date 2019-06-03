## 函数、方法、接口

```go
1. 函数
2. 匿名(闭包)函数
3. 延迟函数(defer)
4. 方法
5. 接口
```

## 1. 函数
> 编写函数的主要目标是将一个复杂问题分解成一系列简单的任务(函数)来解决, 同一个任务(函数)可以被调用多次,有助于代码重用

**函数优势**
> 1. 让程序变得简短清晰
> 2. 有利于程序维护
> 3. 提高代码重用性

**go语言函数特点**
> 1. 函数的参数可以是任意类型(包括函数类型)
> 2. 可以有零个或者多个参数列表
> 3. 可以有零个或者多个返回值列表
> 4. 命名返回值(named return variables),可以给函数的返回值命名
> 5. 可以使用空白标识符忽略返回值
> 6. 参数都是值传递, 也就是函数接收的是传递参数的副本, 如果参数是指针, 指针的值(内存地址)会被复制, 但是指针所指向的地址不会被复制, 我们可以通过这个指针的值来修改这个值所指向的地址的值. 所以引用传递也是按值传递



**注意命名返回值会自动初始化该类型的零值, 如果命名了返回值,我们只需要不带返回值的的return,就完成了返回值**

```go
// 函数声明
// 函数可以有零个或者多个参数列表和零个或者多个返回值列表
func function_name([parameter-list]) ([result-list]) {
}
func function_name([参数列表]) ([返回值列表]) {
}
```
```go
// 示例1: 没有参数没有返回值得函数
// 声明
func f1()  {
  print("hello")
}
// 调用
f1()// 打印: hello

// 示例2: 多参数, 单返回值函数
func sum(a int, b int) int {
  return a + b
}
// 调用
sum(1, 2)// 返回: 3

// 示例3: 多参数, 多返回值的函数
func demo(a string, b string) (string, string) {
	return "a return " + a, "b return " + b
}
// 调用
var a = "aaa"
var b = "bbb"
a1, b1:= demo(a, b)
// a1 = "a return aaa"
// b1 = "b return bbb"
// 示例4: 多参数, 多返回值的函数
_, b1:= demo(a, b)
// b1 = "b return bbb"

// 示例5: 命名返回值函数
// 声明
func nameReturnVariables(t int) (name string, age int) {
	if t == 0 {
		return
	}
	name = "静静"
	age = 18
	return
}
// 调用
name, age := nameReturnVariables(0)
// name = ""
// age = 0

// 调用
name, age := nameReturnVariables(1)
// name = "静静"
// age = 18
```

## 2. 匿名(闭包)函数
>没有函数名,只有函数体的函数

**闭包函数的作用**
> 1. 常用于延迟函数(defer)
> 2. 回调函数和局部封装
> 3. goroutine

```go
// 声明匿名函数并调用
func () {
}()


// 声明匿名函数赋值给变量,后续使用变量名调用
var func1 = func () {
	fmt.Println("匿名函数")
}
func1()
```

**声明匿名函数并调用**
```go
package main

import "fmt"

func main() {
	// 声明匿名函数并调用
	func () {
		fmt.Println("匿名函数")
	}()
}
```

**声明匿名函数赋值给变量,后续使用变量名调用**
```go
package main

import "fmt"

func main() {
	// 声明匿名函数赋值给变量,后续使用变量名调用
	var func1 = func () {
		fmt.Println("匿名函数")
	}
	func1()
}
```


## 3. 延迟(defer)函数
> defer语句用于延迟调用指定的函数, 它只能出现在函数内部.
> defer语句允许我们推迟到函数返回之前(或者任意位置执行return语句之后)一刻才执行某个语句或者函数, (为什么要在返回之后执行这些语句, 因为return语句同样可以包含一些操作, 而不是单纯的返回某个值)

**defer函数特点**
> 1. defer语句的执行顺序和定义顺序相反,即最后定义的最先执行
> 2. defer中的外包变量应该通过参数传入, 否则容易出问题.
> 3. defer所属的函数中定义了命名返回值(named return variables) 在defer中改变了命名返回值的值,就改变了defer函数的返回值
> 4. 当执行defer所属的函数中的return语句时, 所属函数中所有的defer语句执行完毕后, 才会返回.
> 5. 当defer所属的函数宕机时, 所属函数中所有的defer语句执行完毕后, 才会宕机.
```go
// defer语句声明
defer func
defer 函数
```
**示例: defer所属函数无返回值**
```go
package main

import "fmt"

func main() {
	fmt.Println("第1步: main start")
	demoFunc1()
	fmt.Println("第5步: main end")
}

func demoFunc1() {
	fmt.Println("第2步: demoFunc1 start")
	defer func() {
		fmt.Println("第4步: demoFunc1 run defer")
	}()
	fmt.Println("第3步: demoFunc1 end")
}
/*
$ go run main2.go执行结果如下
第1步: main start
第2步: demoFunc1 start
第3步: demoFunc1 end
第4步: demoFunc1 run defer
第5步: main end
*/
```
**示例: defer所属函数有返回值**
```go
package main

import "fmt"

func main() {
	fmt.Println("第1步: main start")
	str := demoFunc1()
	fmt.Println("第5步: main demoFunc1 return value ", str)
	fmt.Println("第6步: main end")
}

func demoFunc1() string {
	str := "demoFunc1.value"
	fmt.Println("第2步: demoFunc1 start", " ", str)
	defer func() {
		str = "demoFunc1.defer.value"
		fmt.Println("第4步: demoFunc1 run defer", " ", str)
	}()
	fmt.Println("第3步: demoFunc1 end", " ", str)
	return str
}
/*
$ go run main2.go
第1步: main start
第2步: demoFunc1 start   demoFunc1.value
第3步: demoFunc1 end   demoFunc1.value
第4步: demoFunc1 run defer   demoFunc1.defer.value
第5步: main demoFunc1 return value  demoFunc1.value
第6步: main end
*/
/*
注意: defer所属函数的返回值不是命名返回值时, defer无法改变返回值
*/
```

**示例: defer所属函数有返回值并且返回值是一个指针**
```go
package main

import "fmt"

func main() {
	fmt.Println("第1步: main start")
	str := demoFunc1()
	fmt.Println("第5步: main demoFunc1 return value ", *str)
	fmt.Println("第6步: main end")
}

func demoFunc1() *string {
	str := "demoFunc1.value"
	fmt.Println("第2步: demoFunc1 start", " ", str)
	defer func() {
		str = "demoFunc1.defer.value"
		fmt.Println("第4步: demoFunc1 run defer", " ", str)
	}()
	fmt.Println("第3步: demoFunc1 end", " ", str)
	return &str
}
/*
$ go run main2.go
第1步: main start
第2步: demoFunc1 start   demoFunc1.value
第3步: demoFunc1 end   demoFunc1.value
第4步: demoFunc1 run defer   demoFunc1.defer.value
第5步: main demoFunc1 return value  demoFunc1.defer.value
第6步: main end
*/
/*
注意: defer所属函数的返回值不是命名返回值时, defer无法改变返回值(除非返回值是一个指针)
*/
```
**示例: defer所属函数有命名返回值**
```go
package main

import "fmt"

func main() {
	fmt.Println("第1步: main start")
	str := demoFunc1()
	fmt.Println("第5步: main demoFunc1 return value ", str)
	fmt.Println("第6步: main end")
}

func demoFunc1() (str string) {
	str = "demoFunc1.value"
	fmt.Println("第2步: demoFunc1 start", " ", str)
	defer func() {
		str = "demoFunc1.defer.value"
		fmt.Println("第4步: demoFunc1 run defer", " ", str)
	}()
	fmt.Println("第3步: demoFunc1 end", " ", str)
	return
}
/*
$ go run main2.go
第1步: main start
第2步: demoFunc1 start   demoFunc1.value
第3步: demoFunc1 end   demoFunc1.value
第4步: demoFunc1 run defer   demoFunc1.defer.value
第5步: main demoFunc1 return value  demoFunc1.defer.value
第6步: main end

*/
/*
注意: defer所属的函数中定义了命名返回值(named return variables) 在defer中改变了命名返回值的值,就改变了defer函数的返回值
*/
```

## 4. 方法
> 方法声明和函数类似, 只是前面多了一个参数, 这个参数把这个方法绑定的对应的类型上.
```go
// 方法声明
func (type) function_name([parameter-list]) ([result-list]) {
}
func (类型) function_name([参数列表]) ([返回值列表]) {
}
```

```go
// 示例1: 自定义类型定义方法
type month int// 定义类型
// 定义方法
func (m month) toString() string {
	return fmt.Sprintf("%d月", m)
}
var m month
m = 1
print(m.toString())// 打印: 1月
m = 2
print(m.toString())// 打印: 2月


// 示例2: 结构体类型定义方法
type User struct {
	Name string
	Age int
}
func (u User) toString() string {
	return fmt.Sprintf("%d的%s", u.Age, u.Name)
}
user := User{
	Name: "静静",
	Age: 18,
}
print(user.toString())// 打印: 18的静静
```
接口泛指实体把自己提供给外界的一种[抽象化](https://baike.baidu.com/item/%E6%8A%BD%E8%B1%A1%E5%8C%96/10844295)物（可以为另一实体），用以由内部操作分离出外部沟通方法，使其能被内部修改而不影响外界其他实体与其交互的方式。

## 5. 接口
> 接口是一种或者多种类似事物(实体)的抽象, 更是一种约束或规范.
> 为什么说接口是一种约束或者规范, 因为调用者和实者均需要遵守的这种协议
> 接口的作用是将定义与实现分离, 降低耦合

**接口特点**
> 1. 包含0个或者多个方法
> 2. 接口可以组合(接口里面可以嵌套接口)
> 3. 接口可以显示声明和隐式声明
> 4. 隐式声明即接口不需要显示的声明它实现了某个接口,即任何类型实现了接口中声明的全部方法,则表明该类型实现了该接口

```
// 接口定义
type interface_name interface{}
type 接口名 interface{}
```

##### 接口示例
> 例如小明要读取usb数据,写以下的代码

```go
package main

import "fmt"

type UsbDriver struct {
	data string
}

// usb获取数据
func (d *UsbDriver) ReadData() string {
	return "usb中的数据: " + d.data
}

// usb写入数据
func (d *UsbDriver) WriteData(data string) {
	d.data = data
}

func main() {
	driver := UsbDriver{}
	PrintDriverData(driver, "18岁的静静")
}

// 打印驱动数据
func PrintDriverData(driver UsbDriver, data string) {
	driver.WriteData(data)
	fmt.Println(driver.ReadData())
}
```
> 小红也需要读取usb数据, 看见小明已经写一份了, 就直接拿过来用, 现在小明usb的数据已经放在电脑中的文件里面, usb现在不需要用了, 这个简单把UsbDriver复制一份,改读取写入方法让后在把PrintDriverData方法的参数driver的类型改成新的就好, 代码如下
```go
package main

import "fmt"

type UsbDriver struct {
	data string
}

// usb获取数据
func (d *UsbDriver) ReadData() string {
	return "usb中的数据: " + d.data
}

// usb写入数据
func (d *UsbDriver) WriteData(data string) {
	d.data = data
}

type FileDriver struct {
	data string
}
// 文件获取数据
func (d *FileDriver) ReadData() string {
	return "文件中的数据: " + d.data
}

// 文件写入数据
func (d *FileDriver) WriteData(data string) {
	d.data = data
}
func main() {
	driver := FileDriver{}
	PrintDriverData(driver, "18岁的静静")
}

// 打印驱动数据
func PrintDriverData(driver FileDriver, data string) {
	driver.WriteData(data)
	fmt.Println(driver.ReadData())
}
```

> 小红突然发现自己的usb读取数据没法用了, 程序编译不通过, 看错误信息, 就给小明说, 我写在用usb读取数据打印数据, 你改了我都没办法编译了. 小明使用接口修改了下代码

```go
package main

import "fmt"

type Driver interface {
	ReadData() string
	WriteData(string)
}

// usb驱动显示的实现了接口
type UsbDriver struct {
	data string
	Driver
}

// usb获取数据
func (d *UsbDriver) ReadData() string {
	return "usb中的数据: " + d.data
}

// usb写入数据
func (d *UsbDriver) WriteData(data string) {
	d.data = data
}

// 文件驱动隐式的实现了接口
type FileDriver struct {
	data string
}
// 文件获取数据
func (d *FileDriver) ReadData() string {
	return "文件中的数据: " + d.data
}

// 文件写入数据
func (d *FileDriver) WriteData(data string) {
	d.data = data
}
func main() {
	driver := &FileDriver{}
	PrintDriverData(driver, "18岁的静静")
}

// 打印驱动数据
func PrintDriverData(driver Driver, data string) {
	driver.WriteData(data)
	fmt.Println(driver.ReadData())
}
```

> 小红和小明都遵守了驱动(Driver)接口的约定或者说规范来使用接口
