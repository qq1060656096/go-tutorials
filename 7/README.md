## 错误和异常处理(error、panic、recover)

> go语言最大限度避免滥用错误和异常, 而滥用异常无论从性能还是可维护性上看都是大忌.虽然这样会使代码显得繁琐.

```go
1. 错误和异常
    1.1 什么是错误
    1.2 什么是异常
    1.3 go语言的错误和异常
2. error
3. panic 触发异常
4. recover 终止异常
```
[参考文章](https://studygolang.com/articles/11753)
[参考文章2](https://www.zhihu.com/question/27158146/answer/44676012)

> 错误和异常需要分类和管理，不能一概而论错误和异常的分类可以以是否终止业务过程作为标准错误是业务过程的一部分，异常不是不要随便捕获异常，更不要随便捕获再重新抛出异常Go语言项目需要把Goroutine分为两类，区别处理异常在捕获到异常时，需要尽可能的保留第一现场的关键数据
```go
1. 错误和异常需要分类和管理，不能一概而论
2. 错误和异常的分类可以以是否终止业务过程作为标准
3. 错误是业务过程的一部分，异常不是
4. 不要随便捕获异常，更不要随便捕获再重新抛出异常
5. Go语言项目需要把Goroutine分为两类，区别处理异常
6. 在捕获到异常时，需要尽可能的保留第一现场的关键数据
```

## 1. 错误和异常
###### 1.1 什么是错误
> 错误是可能出现问题的地方出了问题, 是我们意料之中的问题(比如: 数据库连接失败、打开文件失败或者请求接口失败等)

###### 1.2 什么是异常
> 异常是不应该出问题的地方出现了问题, 是我们意料之外的问题(比如: 引用了空指针)

###### 1.3 go语言的错误和异常
> 1. go语言中引入了error接口类型作为错误处理的标准模式, 如果函数要返回错误,  则返回值列表中肯定包含error类型. error错误处理类似于c语言中的错误码, 可逐层返回, 直到被处理.
> 2. go语言中引入了2个内置函数panic来触发异常, recover来终止异常处理流程, 同时还引入了defer语句延迟处理后面的函数.

## 2. error
> go语言错误: 通过内置的错误接口提供了非常简单的处理处理机制

**go语言error接口定义**
```go
type error interface {
	Error() string
}
```
**错误示例**
```go
package main

import (
	"errors"
	"fmt"
)

func main()  {

	age := 18
	if err := checkAge(age); err == nil {
		fmt.Printf("年龄%d岁合法, err=%#v \n", age, err)
	}

	age = -1
	if err := checkAge(-1); err != nil {
		fmt.Printf("年龄%d岁非法, err=%#v \n", age, err)
	}
}

// 检测年龄是否合法
func checkAge(age int) error {
	if age < 0 || age > 150 {
		return errors.New("年龄非法, 年龄必须在0-150之间")
	}
	return nil
}

/*
$ go run main.go
年龄18岁合法, err=<nil> 
年龄-1岁非法, err=&errors.errorString{s:"年龄非法, 年龄必须在0-150之间"} 

*/
```

## 3. panic 触发异常
> panic触发异常时会触发宕机, 导致程序执行终止

```go
package main

import "fmt"

// 定义月份类型
type month uint8
// 中国月份
var chinaMonth = map[month]string{
	1: "一月",
	2: "二月",
	3: "三月",
	4: "四月",
	5: "五月",
	6: "六月",
	7: "七月",
	8: "八月",
	9: "九月",
	10: "十月",
	11: "十一月",
	12: "十二月",
}

func main()  {
	var m month = 1
	fmt.Printf("获取%d月中文月份: %#v \n", m, getChinaMonth(m))
	m = 13
	// 非法月份, 这里会触发异常
	fmt.Printf("获取%d月中文月份: %#v \n", m, getChinaMonth(m))
}

// 获取中文月份
func getChinaMonth(m month) string {
	if m < 1 || m > 12 {
		panic("月份非法, 1年只有12个月, 1-12之间的数字才有效")
	}
	return chinaMonth[m]
}

/*
$ go run main.go
获取1月中文月份: "一月"
panic: 月份非法, 1年只有12个月, 1-12之间的数字才有效

goroutine 1 [running]:
main.getChinaMonth(...)
        /Users/zhaoweijie/develop/go/tests/main.go:34
main.main()
        /Users/zhaoweijie/develop/go/tests/main.go:28 +0x109
exit status 2
*/
```
![1681559466255_.pic_hd.jpg](https://upload-images.jianshu.io/upload_images/6713312-6f2f1a0f6a92c879.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 4. recover 终止异常
>  recover 终止异常, 会导致应该宕机并执行终止程序, 正常执行. (当异常需要转错误时,我们需要终止异常)

```go
package main

import "fmt"

// 定义月份类型
type month uint8
// 中国月份
var chinaMonth = map[month]string{
	1: "一月",
	2: "二月",
	3: "三月",
	4: "四月",
	5: "五月",
	6: "六月",
	7: "七月",
	8: "八月",
	9: "九月",
	10: "十月",
	11: "十一月",
	12: "十二月",
}

func main()  {
	var m month = 1
	fmt.Printf("获取%d月中文月份: %#v \n", m, getChinaMonth(m))
	m = 13
	// 非法月份, 这里会触发异常
	fmt.Printf("获取%d月中文月份: %#v \n", m, getChinaMonth(m))
}

// 获取中文月份
func getChinaMonth(m month) string {
	// 函数返回之前, return语句之后, 会调用defer recoverFunc方法, 该方法会终止异常发生
	defer recoverFunc()
	if m < 1 || m > 12 {
		panic("月份非法, 1年只有12个月, 1-12之间的数字才有效")
	}
	return chinaMonth[m]
}

// 终止异常
func recoverFunc() {
	if err := recover(); err != nil {
		fmt.Printf("recoverFunc.stop.panic, err=%#v \n", err)
	}
}
/*
go run main.go
获取1月中文月份: "一月"
recoverFunc.stop.panic, err="月份非法, 1年只有12个月, 1-12之间的数字才有效"
获取13月中文月份: ""
*/
```
![1691559467319_.pic_hd.jpg](https://upload-images.jianshu.io/upload_images/6713312-c5fd43efff304c82.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
