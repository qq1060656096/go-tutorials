## 测试
> go语言拥有一套单元测试、性能测试系统和示例测试，仅需要添加很少的代码就可以快速测试一段需求代码
> testing 包提供自动测试支持, 它与"go test"命令一起使用, 该命令自动执行*_test.go源文件中以Test、Benchmark和Example前缀的函数

**为什么要些测试？**

[更多如何写好单元测试文章](https://www.jianshu.com/p/1cb94e6508b8)
```sh
目的：
1. 提高软件质量
2. 减少bug
3. 减少重复的工作
4. 安全的重构已有的代码
5. 让开发者对程序稳定性更有信心

重要性：
1. 运行单元测试是为了保证代码的行为和我们期望的结果一致。
2. 写单元测试会增加代码工作量,同时也节约了bug修复时间。
3. 如果没有写单元测试，没有发现bug的情况下，程序在测试人员测试的时候才发现问题或者在线上环境(正式环境)用户使用才发现问题，在去修复bug。开发会花大量的精力去修复bug和走部署流程，测试也会花大量的时间去做了重复的测试。很不划算。
4. 在线上某些场景下bug导致大量的数据丢失,需要花很大精力去修复数据，或者根本没办修复数据导致严重的后果。
```


```go
1. Test功能测试函数
2. Benchmark基准测试函数
3. Example示例函数
4. 使用流行开源的测试包(github.com/stretchr/testify)
5. 白盒测试
6. 代码覆盖率
7. 性能剖析
```

## 1. Test功能测试函数

> 1. 每个测试文件都必须导入testing包
> 2. 功能测试函数必须以Test为开头(前缀)

```go
// 文件: *_test.go
func TestXxx(t *testing.T) {
    // 执行代码
    ...
}
```

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
![功能测试结果](images/examples.demo1.run.test.result.jpg)

## 2. Benchmark 基准测试函数
> 1. 基准测试可以测试一段程序在给定的工作负载下检测程序性能的一种方法, 可以测试一段程序的运行性能及耗费CPU的成都.
> 2. go语言提供了基准测试框架,使用方式和功能测试类似
> 3. 使用者无需准备高精度的计时器和各种分析工具, go语言的基准测试本身就可以打印出非常标准的测试报告.

```go
// 文件: *_test.go
// 基准测试声明
func BenchmarkXxx(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // 执行代码
    }
}
```

**基准测试示例**
> 代码请看本目录下examples/demo2包
```sh
# 进入目录
cd examples/demo2
# 运行所有基准测试
go test -bench=. -v

```
![基准测试结果](images/examples.demo2.run.bench.result.jpg)

## 3. Example示例函数
> Example示例函数主要目的作为文档, 与带注释的示例不同的是, Example示例函数是真实的go语言代码, 必须通过编译时检查,所以随着代码的进化它们也不会过时

```go
// 文件: example_test.go
// Example示例函数声明
func ExampleXxx() {
	// 执行代码
	...
}
```

**Example示例函数示例**
> 代码请看本目录下examples/demo3包
```sh
# 进入目录
cd examples/demo3
# 运行所有基准测试
go test -v

```
![示例函数执行结果](images/examples.demo3.run.example.result.jpg)

## 4. 使用流行开源的测试包(github.com/stretchr/testify)

**特点**
> 1. 快速断言(assertions)
> 2. 打桩(mocking)
> 3. 套件(suite)和功能(function)

**示例**
> 代码请看本目录下examples/demo4包
```sh
# 进入目录
cd examples/demo4
# 运行所有基准测试
go test -v

```
![第三方测试包断言结果](images/examples.demo4.other_test.result.jpg)


## 5. 白盒测试
> 1. 白盒测试又称为 结构测试、透明测试、逻辑测试或基于代码的是. 白盒指盒子是透明的, 我们可以清楚的知道盒子里面的一切.
> 2. 白盒测试要考虑到测试用例对程序内部的覆盖程度, 最好能覆盖到每一条路径.
> 3. 白盒测试可对实现的特定之处提供更加详细的覆盖测试
> 4. 黑盒测试就当整个程序是个黑盒子, 我们看不到它里面做了些什么事情, 只能通过公开api和文档输入输出来测试
> 5. 黑盒测试更加健壮，每次程序更新后基本不需要修改

**白盒测试示例**
> 1. 创建examples/demo5/login.go
> 2. 创建examples/demo5/login_test.go
> 3. 运行测试

> 1. 创建examples/demo5/login.go并添加一下内容
```go
// examples/demo5/login.go
package demo5

import (
	"fmt"
	"log"
	"net/smtp"
)

const SenderUser  = "1060656096@qq.com"
const SenderPass = "123456"
const Hostname = "smtp.qq.com"

// 发送邮件
var SendLoginEmail = func (user string) {
	msg := fmt.Sprintf("%s, Welcome to login", user)
	auth := smtp.PlainAuth("", SenderUser, SenderPass, Hostname)
	if err := smtp.SendMail(Hostname, auth, SenderUser, []string{user}, []byte(msg)); err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", user, err)
	}
}

// Login 登录
func Login(user, pass string) bool {
	if user == "root" && pass == "123456" {
		SendLoginEmail(user)
		return true
	}
	return false
}

```

> 2. 创建examples/demo5/login_test.go并添加一下内容
```go
package demo5

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestLogin(t *testing.T) {
	// 保留恢复SendLoginEmail, 避免后续测试混乱
	var oldSendLoginEmail = SendLoginEmail
	defer func() {
		SendLoginEmail = oldSendLoginEmail
	}()

	SendLoginEmail = func(user string) {
		fmt.Println(user)
		assert.Equal(t, "root", user)
	}

	b := Login("root", "123456")
	assert.Equal(t, true, b)
}
```

> 3. 运行测试
```sh
# 进入目录
cd examples/demo5
# 运行测试
go test -v

```
![白盒测试结果](images/examples.demo5.test.result.jpg)

## 6. 代码覆盖率
> 1. 从本质上来讲, 测试是发现bug, 而不是证明没有bug, 无论有多少个测试都无法证明是没有bug的.
> 2. 就算覆盖率达到100%，和"靠谱"的代码肯能有天壤之别，问题就在于有些公司把代码覆盖率作为考核的标准，这就让开发很容易就演变成"追求100%代码覆盖率"，然后无所不用，那就更悲剧了，一群人对着水分极大的代码，然后对着"100%代码覆盖率"乐得合不弄嘴，想想都难受想哭。

```sh
# 进入目录
cd examples/demo5
# 运行测试
go test -coverprofile=c.out -v
go test -coverprofile=c.out
```
**代码覆盖率测试结果**
```sh
$ go test -coverprofile=c.out
PASS
coverage: 37.5% of statements
ok      business-card/docker-compose/go-tutorials/8/examples/demo5      0.014s
```

## 7. 性能剖析

```sh
# 进入目录
cd examples/demo5
# 运行测试
go test -cpuprofile=c.out && go tool pprof -text c.out
# 执行500次测试的cpu性能分析
go test -cpuprofile=c.out -count 500 && go tool pprof -text c.out
go test --memprofile=mem.out -count 1000 && go tool pprof -text mem.out
go test --blockprofile=block.out -count 500 && go tool pprof -text block.out

```
**性能剖析测试结果**
```sh
# 执行500次测试的cpu性能分析
$ go test -cpuprofile=c.out -count 500 && go tool pprof -text c.out
PASS
ok      business-card/docker-compose/go-tutorials/8/examples/demo5      0.239s

Type: cpu
Time: Jun 5, 2019 at 10:14pm (CST)
Duration: 205.64ms, Total samples = 10ms ( 4.86%)
Showing nodes accounting for 10ms, 100% of 10ms total
      flat  flat%   sum%        cum   cum%
      10ms   100%   100%       10ms   100%  runtime.pthread_cond_signal
         0     0%   100%       10ms   100%  runtime.mstart
         0     0%   100%       10ms   100%  runtime.newproc.func1
         0     0%   100%       10ms   100%  runtime.newproc1
         0     0%   100%       10ms   100%  runtime.notewakeup
         0     0%   100%       10ms   100%  runtime.semawakeup
         0     0%   100%       10ms   100%  runtime.startm
         0     0%   100%       10ms   100%  runtime.systemstack
         0     0%   100%       10ms   100%  runtime.wakep
```