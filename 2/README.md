## 2. 变量、常量、指针

```
1. 变量
2. 常量
3. 指针
```
## 1. 变量
> 变量有3种声明方法
```
1. 第1种,指定变量类型,声明后不赋值,使用默认值
2. 第2种,根据值类型自动判定变量类型
3. 第3种,":="短变量声明(注意短变量声明最少声明一个新的变量)
4. 全局变量声明
```

> 变量类型对应默认值
````
对于数字值是0
对于布尔值是false
对于字符串是""(即空字符串)
对于接口和引用类型(slice, 指针, map, 通道, 函数)值是nil
````
###### 1.1 第1种,指定变量类型，声明后不赋值，使用默认值
```go
// var name   type
// var 变量名 变量类型
var str string //string=""

var id, age int // int=0, int=0
```

###### 1.2. 第2种,根据值类型自动判定变量类型 

```go
// 变量声明
// var name   type     = expression
// var 变量名 变量类型 = 表达
var str string = "str value"

var str2 = "str2 value"

var str3
str3 = "str3 value"

//  多重赋值
var id, age int = 1, 10 // int=1, int=10)

var i, f, b = 0, 1.0, true // int=0, float64=1, bool=true)

var id, age int
id, age = 1, 20 // int=1, int=20)
```

###### 1.3. 第3种,":="短变量声明
> ***注意短变量声明最少声明一个新的变量***  
> ***短声明变量一般用在局部变量***  
```go
// name   := expression
// 变量名 := 表达式

id := 10 // int=10

id, name, hasMenu:= 10, "张三", true // int=10, string=张三, bool=true


// 以下声明会编译错误
id, name, hasMenu:= 10, "张三", true
id, name, hasMenu:= 10, "张三", true
```

**编译错误的短变量声明**
```go
id, name, hasMenu := 10, "张三", true
id, name, hasMenu := 10, "张三", true // 编译错误: 没有新的变量
```

**正确的短变量声明**  
> ***短变量声明最少声明一个新的变量***
```go
id, name, hasMenu := 10, "张三", true // int=10, string=张三, bool=true
id, name, hasMenu, sex := 10, "张三", true, "男" // int=10, string=张三, bool=true, string=男
```

###### 1.4. 全局变量声明
```go
var (
    id int = 1 
    name string = "张三"
)

var (
    id = 1 // int=1
    name = "张三" // string=张三
)

var (
    id int // int=0
    name string // string=
)
```

## 2. 常量
```
// 常量声明
// const name   type     = expression
// const 变量名 变量类型 = 表达式
const id int = 1
const id = 1
const (
    id int = 1
    name string = "张三"
)

const (
    id = 1 // int=1
    name = "张三" // string=张三
)
```

## 3. 指针
> 3.1 什么是指针  
> 3.2 指针限制  
> 3.3 指针声明  
> 3.4 取地址操作符&(***唯一可以取地址的操作符***)


###### 3.1 什么是指针

```
指针的值是一个变量的地址. 一个指针指向的值所保存的位置.  
不是所有的值都有地址, 但是所有的变量都有.
使用指针, 可以在无须知道变量名字的情况下, 间接读取或者更改变量值.
```

###### 3.2 指针的限制

```
1. 不同类型的指针不能互相转化, 例如*int, int32, 以及int64
2. 任何普通指针类型*T和uintptr之间不能互相转化
3. 指针变量不能进行运算, 比如C/C++里面的++, --运算
```

###### 3.3 指针声明

```go
// 指针声明
// var name   * type
// var 变量名 * 类型

var p * int // *int=<nil>


var i = 1
var p * int // 声明指针p
p = &i // 指针指向变量地址(即在指针中存放变量地址): *int=0xc042054080
*p = 10 // 间接变量,设置指针指向变量的值：
fmt.Printf("p=%v, *p=%v, i=%v", p, *p, i) // p=0xc042054080, *p=10, i=10


// 指针判断
var p * int // p=<nil>
if (p == nil) { // 是空指针
    fmt.Printf("p=%v", p) // p=<nil>
}
i := 10
p = &i
if (p != nil) { // 不是空指针
    fmt.Printf("p=%v", p) // p=0xc0420100d0
}
```