## 测试
> go语言 testing 包提供自动测试支持, 它与"go test"命令一起使用, 该命令自动执行*_test.go源文件中以Test或者Benchmark前缀的函数


```go
1. Test功能测试函数
2. Benchmark基准测试函数
3. 示例函数

```

###### 1. Test功能测试函数
> 1. 每个测试文件都必须导入testing包
> 2. 功能测试函数必须以Test为开头(前缀)

**功能测试示例**
> 1. 创建examples/demo1/login.go
> 2. 创建examples/demo1/login_test.go
> 3. 运行测试

> 1. 创建examples/demo1/login.go并添加一下内容
```go
// examples/demo1/login.go
package demo1

// Login 登录
func Login(user, pass string) bool {
	if user == "root" && pass == "123456" {
		return true
	}
	return false
}
```
> 2. 创建examples/demo1/login_test.go并添加一下内容

```go
// examples/demo1/login_test.go
package demo1

import "testing"

// 成功的测试
func TestLoginSuccess(t *testing.T) {
	user := "root"
	pass := "123456"
	isLogin := Login(user, pass)
	// 登录失败
	if !isLogin {
		t.Errorf("user=%s,pass=%s, 用户名必须是root,密码必须是123456", user, pass)
	}
}

// 失败的测试
func TestLoginFail(t *testing.T) {
	user := "admin"
	pass := "123456"
	isLogin := Login(user, pass)
	// 登录失败
	if !isLogin {
		t.Errorf("user=%s,pass=%s, 用户名必须是root,密码必须是123456", user, pass)
	}
}
```
> 3. 运行测试
```sh
# 进入目录
cd examples/demo1
# 运行测试
go test -v
```
![功能测试结果](./images/examples.demo1.run.test.result.jpg)