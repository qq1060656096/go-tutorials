## 3. 数据类型

```go
1. 基本数据类型
    1.1 整数
    1.2 浮点
    1.3 复数
    1.4 布尔（boolean）
    1.5 字符串（string）
	1.6 rune unicode码点（UTF-8字符串中的字符）
2. 复合数据类型
    2.1 数组
    2.2 切片(slice)
    2.3 map
    2.4 结构体
3. 任意数据类型(空接口)
    interface{}
4. 类型断言
    type.(类型)
5. 类型转换
    type()
```
## 1. 基本数据类型
###### 1. 1 整数
> 有符号和无符号分别各自对应四种大小(8位、16位、32位、64位)
> 注意int,uint在特定的平台大小相等, 在该平台运行效率最高, 即使在相同硬件平台上, 不同的编译器可能选用的大小不同
```go
  1. 无符号整数: uint, uint8, uint16, uint32, uint64
  2. 有符号整数: int, int8, int16, int32, int64
  3. 无符号整数uintptr: 大小并不明确, 但足够存放任何指针
// 声明
var id uint64 = 1
```

###### 1. 2 浮点
```go
浮点只有2种: float32, float64
// 声明
var f float32 = 10.5
```
###### 1. 3 复数
```go
复数只有2种: complex64, complex128
// 声明
var f complex64 = 1 + 2i
```
###### 1. 4 布尔(bool)
```go
# 布尔值只有2中可能: 真(true), 假(false)
// 声明
var yes bool = true
```
###### 1. 5 字符串(string)
> 字符串就是一串固定长度的字符连接起来的字符序列. go语言字符串是由单个字节连接起来的.go语言的字符串的字节使用UTF-8编码标识unicode文本.是不可变的字符序列, 意味着不能修改字符串中的字符, 但是可以通过构造新的字符串赋值给原来的字符串变量来改变. 字符串及其子字符串可以安全地共用数据,因此生成子字符串的操作开销低廉.

> **go语言字符串特点**
> 1. 由单个字节(byte)组成
> 2. **是不可变的字节序列**
> 3. 生成子字符串的操作开销低廉
```go
var s string = "hello 你好"
//  字符串转字节序列
b := []byte(s)
// 字符串转unicode码点，int32 别名，用于区分数字和字符
runes := []rune // []int32{'h', 'e', 'l',' l', 'o', ' ', '你', '好'}

```

## 2. 复合数据类型

###### 2.1 数组
> 数组是具有固定长度拥有0个或者多个相同类型元素的序列, 由于数组长度固定, 所以go语言很少使用数组
```go
var 变量名 [长度]类型
var arr [2]int// [0 0]
arr1 := [...]int{0,1}// [0 1]
arr2 := [2]int{0,1}// [0 1]
```
###### 2.2  切片(slice)
> slice是具相同类型元素的可变长度序列. 它像一个没有长度的数组类型
> **增加切片值: append(slice, 值 ...Type) []Type**
> **删除切片值:**
```go
// slice默认值: nil
var 变量名 []类型
var arr []int// nil
arr1 := []int{0,1}// [0 1]
```
###### 2.3 map
> map一种无序的键值对的集合, map可以通过key快速的检索数据, key类似索引, 指向数据的值
> map是无序的,这是因为map是使用 hash 表来实现的. 我们无法通过for循环有序的处理, 因为每次for循环map的顺序都不一样

```go
// map默认值: nil
var 变量名 map[类型]类型
var users map[string]string// map[string]string(nil)
users1 := map[string]string{}// map[string]string{}
// 通过键访问值, 如果键不存在,将得到map值类型的零值
value:= users1["one"]
// 检测map中是否存在指定的键
value, ok := users1["one"]// 获取值, ok返回键是否存在
if ok {
	// 存在
} else {
	// 不存在
}

// 删除map中的键, 及时键不存在, 删除操作也是安全的
delete(users1, "one")
```
###### 2.4 结构体
> 结构体是将零个或者多个任意类型的命名变量组合在一起的聚合数据类型. 每个变量都叫结构体的成员
```go
// 默认值是类型的零值
// 定义结构体
type 结构体名 struct {
    成员名1: 类型
    成员名2: 类型
    成员名n: 类型
    ...
}
// 访问结构体成员: 结构体.成员名

type User struct {
	Name string
	Age uint8
}
var u User // User{Name:"", Age:0}
var up *User// nil
// 访问结构体成员
u.Name="静静"// 静静
u.Age = 18// 18

// 结构体组合
type UserEmail struct {
	User
	Email string
}
var userEmail UserEmail
userEmail.Name = "小微"
userEmail.Age = 19
userEmail.Email = "wmail@klsh.com"
userEmail// UserEmail{User:User{Name:"小微", Age:19}, Email:"wmail@klsh.com"}
```
## 3. 任意数据类型(空接口)
> interface{}任意类型(空接口)变量可赋值任意类型值
> 空接口没有任何方法, 因此任何类型都无须实现空接口. 从实现的角度看, 任何值都满足这个接口的需求.因此空接口类型可以保存任何值, 也可以从空接口中取出原值
```go
var any interface{}
any = "静静今年18岁"// 赋值string
any = 18// 赋值int
any = false// 赋值布尔
```

## 4. 类型断言
> 类型断言用于检测和转换接口变量的类型
> 注意: 类型断言失败会导致操作崩溃(错误发生并宕机)
> 断言变量必须是接口类型
```go
v := var_name.(type)
值 := 变量名.(类型名)
// 断言字符串
var str interface{}
str = "18岁的静静"
value := str.(string)
// 检测是否断言成功
value, ok := str.(string)
if ok {
	// 断言成功
} else {
	// 断言失败
}
```

## 5. 类型转换
> 类型转换是将一种数据类型的变量转为另一种类型的变量
```go
type(var_name)
类型名(变量名)
var money float32 = 20.5
i := int(money)// 浮点转为整型
```
