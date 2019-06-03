## 4. 控制语句if、switch、for

```go
1. if语句
2. switch语句
3. for语句
```

```go
break语句: 经常用于中断当前 for 循环或跳出 switch 语句 |
continue语句: 跳过当前循环的剩余语句,然后继续进行下一轮循环。
goto 语句: 将控制转移到被标记的语句.
```

## 1. if语句
> if语句用于检测某个条件(布尔bool或者逻辑性)的语句, 如果条件成立, 就会执行 if后面大括号内的代码块

> if语句没有数量限制, 为了代码可读性,尽可能先满足的条件放到前面
```go
if 条件 {
  // 条件true执行的代码
}

if 条件 {
  // 条件true执行的代码
} else {
  // 条件false执行
}

if 条件1 {
  // 条件1true执行的代码
} else if  条件2 {
  // 条件1 false并且条件2 true执行
} else {
  // 以上条件都不满足是执行
}

// 示例1: 布尔条件
if  true {
}
// 示例2: 逻辑条件
if  str == "" {
}

// 示例3: 短声明赋值逻辑条件
// 注意声明的变量只存在if作用域内
if age := 10; age > 18 {
}
// 类似于这种
{
  var age = 10
  if age > 18 {
  }
}

```





## 2. switch语句
> switch语句中所有的switch表达式和case表达式都会被求值
> 并且执行顺序是从左到右,自上而下, 直到匹配到某个case或者default为止
> 一旦匹配到某个条件, 就退出整个switch块, 所以说你不要使用break
```go
// 变量var1可以是任何类型, case语句必须是相同类型
switch var1 {
    case val1:
        ...
    case val2:
    default:
}
// 您可以同时测试多个可能符合条件的值, 使用逗号分割它们,例如：case val1, val2, val3。
switch var1 := 2; var1 {
case 1:// 单值
	fmt.Println("case ", var1)
case 2,3,4,5:// 多值
	fmt.Println("case 2-5 ", var1)
default:
	fmt.Println("default ", var1)
}

// 逻辑条件匹配
switch {
    case condition1:
        ...
    case condition2:
... default:
... }

var age uint64
在age = 18
switch {
case age < 18:
	fmt.Println("age < 18 ", age)
case age == 18:
	fmt.Println("age = 18 ", age)
	fallthrough// 强制执行后面的语句
default:
	fmt.Println("age > 18 ", age)
}
```

## 3. for语句
> for语句用于执行重复的代码

**for range注意事项**
> 1. range 值只会被求值一次
> 2. value值始终是集合对应值得拷贝
> 3. 如果value值对应应集合值是指针,会产生指针拷贝
> 4. 迭代nil的通道值上,会让当前for循环永远阻塞(原因range的值只会被求值一次)
> 5. 安全的unicode字符串迭代
```go
for [condition |  ( init; condition; increment ) | Range] {
}
for 初始化;条件;修饰语句(如i++) {
}

// for range可以对字符串 (string)、数组 (array)、切片 (slice)、map、通道(chan)进行迭代循环
for key, value := range [ 字符串 string | 数组 array | 切片 slice | map | chan] {
}
for key := range [ 字符串 string | 数组 array | 切片 slice | map | chan] {
}
```
```go

// 示例1: 无限循环
var i = 0
for {
	fmt.Println(i)
	i ++
}
// 示例2: 带条件无限循环
var i = 0
for true {
	fmt.Println("true", i)
	i ++
}

// 示例3: 直接使用分号省略初始化、条件、修饰语句的无限循环
for ;; {
	fmt.Println(";; ", i)
	i ++
}

// 示例4: 带始化、条件、修饰语句的for
for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// 示例5: 省略初始化的for
var i = 0
for ; i < 10; i++ {
  fmt.Println(i)
}

// 示例5: 没有修饰语句的for
for i := 0; i < 10; {
	fmt.Println(i)
	i++
}
// 示例6: 没有修饰语句的for,如果没有最后一行左括号不能放在同一行,
// 这是go语言中左括号都是在语句开始的同一行,但是这里例外
// 左花括号放在同一行会提示预发错误
for i := 0; i < 10
{
	fmt.Println(i)
	i++
}

// 示例7: 只带条件的for
var i = 0
for i < 10 {
	fmt.Println(i)
	i ++
}
// 示例8: for range带key,value
var str = "go语言go"
for key, value := range str {
	fmt.Printf("key=%#v, value=%q\n", key, value)
}
/*
key=0, value='g'
key=1, value='o'
key=2, value='语'
key=5, value='言'
key=8, value='g'
key=9, value='o'
*/

//示例9: for range带key
var str = "go语言go"
for key := range str {
	fmt.Printf("key=%#v\n", key)
}
/*
key=0
key=1
key=2
key=5
key=8
key=9
*/
```