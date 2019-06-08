## go语言 包(mod)管理
> Go语言从诞生之初就一直有个被诟病的问题: 缺少一个行之有效的“官方”包依赖管理工具. 其原因是在Google内部，所有人都是在一个代码库上进行开发的,因此并不是非常需要.但是Go语言变成一个社区化的工程语言之后,这个问题被放大了.

```
1. 模块是相关Go包的集合(即一个模块可以包含多个package,一个包package包含多个go源文件)
2. go命令直接支持使用模块
3. 模块中记录和解决对其他模块的依赖性
4. 模块取代了旧的基于GOPATH的方法来指定
5. 有利于程序维护
6. 提高代码重用性(供他人使用)
7. 多版本共存(即同时使用同一个模块多个版本,例如为了更好的升级模块,我们先修改1小部分代码用新的版本,当模块新版本稳定后,我们在全面升级)
```

### 1. 设置 GO111MODULE
```sh
可以用环境变量 GO111MODULE 开启或关闭模块支持，它有三个可选值：off、on、auto,默认值是 auto.
1. GO111MODULE=off 无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。
2. GO111MODULE=on 模块支持，go 会忽略 GOPATH 和 vendor 文件夹,只根据 go.mod 下载依赖.
3. GO111MODULE=auto 在 GOPATH/src 外面且根目录有 go.mod 文件时,开启模块支持.

```

### 2. 如何使用go模块
```sh
1. 把项目放到$GOPATH/src之外
2. 在项目目录下创建模块: "go mod init 模块名",创建模块后,会在模块所在的文件夹生成go.mod文件
3. 然后在项目目录下运行命令: "go build" 、"go test" 或 "go run"执行时，会自己去修改go.mod文件，生成"go.sum"文件
```

**go模块示例**
> 1. 创建模块目录并进入: mkdir examples/hellomod && cd examples/hellomod
> 2. 创建模块: go mod init "hellomod"
> 3. 创建模块
> 3. 提交代码到github上

> 2. 创建模块 hellomod: go mod init "hellomod"
```sh
# 1. 在github上创建仓库 hellomod
# 为什么要创建创库, 为了其他人也可以使用这个模块

# 2. 进入examples目录
cd examples 

# 3. 下载 hellomod 仓库
git clone git@github.com:qq1060656096/hellomod.git

# 4. 进入 hellomod目录
cd hellomod

# 5. 创建模块
go mod init "github.com/qq1060656096/hellomod"
# 创建模块失败会提示: "go: modules disabled inside GOPATH/src by GO111MODULE=auto; see 'go help modules'"

# 为什么创建模块失败
# 因为GO111MODULE默认值是auto, 在 GOPATH/src 之外的目录才开启模块支持
# 我们有2中方式解决以上问题
#   第1种: 在 GOPATH/src 之外的目录创建模块
#   第2种: 直接设置GO111MODULE=on 模块支持

# 这里我们直接使用第2种
export GO111MODULE=on
go mod init "github.com/qq1060656096/hellomod"
# 创建模块成功会提示"go: creating new go.mod: module github.com/qq1060656096/hellomod"
# 模块创建后里面会有一个go.mod文件

# 6. 查看go.mod文件的内容
$ cat go.mod
module github.com/qq1060656096/hellomod
# 里面只有一行, 就定义的模块名字
```

> hellomod模块目录下,创建hello.go文件, 并增加以下内容
```go
package hellomod

func Hello() string {
	return "Hello World!"
}
```

> hellomod模块目录下,创建hello_test.go文件, 并增加以下内容
```go
package hellomod

import "testing"

func TestHello(t *testing.T) {
	want := "Hello World!"
	if Hello() != want {
		t.Errorf("Hello() != %s", want)
	}
}
```

> hellomod模块目录下,运行模块测试 "go test -v"会输出以下内容
```sh
=== RUN   TestHello
--- PASS: TestHello (0.00s)
PASS
ok      github.com/qq1060656096/hellomod        0.004s
```
> 提交 hellomod 模块代码到github
```sh
# 提交代码到github
git add .
git commit -m 'aaa: go hello模块第一次提交'
git push origin master
```

# 回到examples目录并创建一个模块测试 testmod
```sh
# 回到 examples 并创建 testmod 目录, 然后在进入 testmod 目录
cd ../ && mkdir testmod && cd testmod
# 创建 testmod 模块
go mod init "testmod"

```

### testmod 模块中创建 main.go 文件

```go
package main

import (
	"fmt"
	"github.com/qq1060656096/hellomod"
)

func main()  {
	fmt.Println(hellomod.Hello())
}
```

### testmod 模块中执行命令: go run main.go

```sh

# go run main.go
# 命令会输出以下内容
go: finding github.com/qq1060656096/hellomod v0.0.1
go: downloading github.com/qq1060656096/hellomod v0.0.1
Hello World!
```