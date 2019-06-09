## go 常用命令

```sh
1. build 命令 编译包和依赖项
2. clean 命令 删除对象文件和缓存的文件
3. doc与godoc 命令 显示包文档
4. env 命令 打印Go语言的环境信息
5. fix与go tool fix 命令 会把指定包中的所有Go语言源码文件中旧版本代码修正为新版本的代码,升级版本时非常有用
6. fmt与gofmt 命令 格式化go源文件
7. generate 命令
8. get 命令 下载并安装包和依赖(下载包和依赖,并对它们进行编译安装)
9. install 命令 编译并安装指定包及它们的依赖包,
10. list 命令 列出包和模块信息
11. mod 命令
12. run 命令
13. test 命令
14. tool 命令
15. version 命令
16. vet 命令
```

## 1. build 命令
```sh
build 编译包和依赖项
usage: go build [-o output] [-i] [build flags] [packages]
 用法:  go build [-o output] [-i] [build 参数] [包]
 可选参数:
    -o 编译单个包才能使用(不能同时对多个代码包进行编译),例如我们经常重新命名可执行文件名字
    -i 标志安装的包是目标的依赖项
    -a 强制重新构建已经更新的包
    -n 打印编译时执行的命令,但不真正执行
    -p 开启并发编译,默认情况下该值为CPU逻辑核数
    -race 开启竞态检测,只支持linux/amd64、freebsd/amd64、darwin/amd64和windows/amd64.
    -msan 内存扫描
    -v 编译包时打印包的名称
    -work 编译时打印临时工作目录的名称,退出时不删除它
    -x 打印编译时执行的命令(打印编译时会用到的所有命令)
    -asmflags '[pattern=]arg list',传递给每个go工具asm调用的参数
    -buildmode mode 使用编译模式, 更多信息请参见“go help buildmode”
    -compiler name 设置编译时使用编译器名,编译器名称只有2个选项(gccgo或gc)
    -gccgoflags '[pattern=]arg list' 传递每个gccgo编译器/链接器调用的参数列表
    -gcflags '[pattern=]arg list' 用于指定需要传递给go tool compile命令的参数的列表,更多信息参见(go tool compile)
    -installsuffix suffix 为了使当前的输出目录与默认的编译输出目录分离,可以使用这个标记.此标记的值会作为结果文件的父目录名称的后缀.其实,如果使用了-race标记,这个标记会被自动追加且其值会为race.如果我们同时使用了-race标记和-installsuffix,那么在-installsuffix标记的值的后面会再被追加_race,并以此来作为实际使用的后缀
    -ldflags '[pattern=]arg list' 用于指定需要传递给go tool link命令的参数的列表
    -linkshared
    -mod mode 模块下载方式,只有2个选项(readonly或vendor),更多信息请参见(go help modules)
    -pkgdir dir 设置包目录.编译器会只从该目录中加载代码包的归档文件,并会把编译可能会生成的代码包归档文件放置在该目录下
    -tags 'tag list'
    -toolexec 'cmd args' 用于在编译期间使用一些Go语言自带工具(如vet、asm等)的方式
示例:
go build [1个或多个go源文件, 或者包名, 或者包所在目录]
go build a.go b.go main.go
go build main.go
go build hello
# 把main.go编译成可执行文件hello.exe
go build -o hello.exe  main.go
# 打印编译时执行的命令,但不真正执行
go build -n 
# 答应工作目录
go build -work 
```

## 2. clean 命令
```sh
clean 删除执行其他命令时产生的文件、目录和缓存文件.
    具体地说.clean 会删除从导入路径对应的源码目录中,删除以下这些文件和目录
        _obj/            old object directory, left from Makefiles
        _test/           old test directory, left from Makefiles
        _testmain.go     old gotest file, left from Makefiles
        test.out         old test log, left from Makefiles
        build.out        old test log, left from Makefiles
        *.[568ao]        object files, left from Makefiles
        DIR(.exe)        from go build
        DIR.test(.exe)   from go test -c
        MAINFILE(.exe)   from go build MAINFILE.go
        *.so             from SWIG
usage: go clean [clean flags] [build flags] [packages]
  用法: go clean [clean 参数]   [build参数]   包
  可选参数:
    -i 会删除安装当前代码包所有产生的所有文件, 如果当前包是一个普通包(不是main包),则结果文件指的就是在工作区的pkg目录的相应目录下的归档文件.如果当前代码包中只包含一个命令源码文件, 则删除当前目录和在工作区的bin目录下的可执行文件和测试文件.
    -n 打印clean执行的命令,但不真正执行
    -r 删除当前代码包和所有的依赖包产生的文件、目录和缓存文件
    -x 打印clean执行的删除命令
    -cache 删除所有 go build 的缓存
    -testcache 删除当前包所有的测试结果
```

## 3. doc 命令
```sh
doc与godoc 显示包或符号的文档, 更多用法请参考(godoc -h)
usage: go doc [-u] [-c] [package|[package.]symbol[.methodOrField]]
 用法:  go doc [-u] [-c] [package|[package.]symbol[.methodOrField]]
 可选参数:
    -c 区分参数包名的大小写.默认情况下,包名是大小写不敏感的
    -cmd 打印 main 包文档, 默认情况下不会打印 main 包文档
    -u 打印出所有的文档(同事包含可导出和不可导出实体)

示例:
# 显示 hellomod 包文档,(注意 hellomod 和 Hellomod是不同的包)
go doc -c hellomod
```

## 4. env 命令

```sh
env 打印Go语言的环境信息
usage: go env [-json] [var ...]
 用法: go env [-json] [变量 ...]
 可选参数:
    -json 以json格式打印环境信息

示例:
# 以json格式打印所有环境信息
go env -json 

# 以json格式只打印 GOOS 程序构建环境的目标操作系统
go env -json GOOS

# 打印所有环境信息
go env

# 只打印 GOOS 程序构建环境的目标操作系统
go env GOOS
```

## 5. fix与go tool fix 命令
```sh
fix 会把指定包中的所有Go语言源码文件中旧版本代码修正为新版本的代码
usage: go fix [packages]

示例:
go fix testmod


go tool fix -h
usage: go tool fix [-diff] [-r fixname,...] [-force fixname,...] [path ...]
    -diff 不将修正后的内容写入文件, 而只打印修正前后的内容的对比信息到标准输出
    -force string 使用此参数后, 即使源码文件中的代码已经与Go语言的最新版本相匹配, 也会强行执行指定的修正操作.该参数值就是需要强行执行的修正操作的名称,多个名称之间用英文半角逗号分隔
    -r string 只对目标源码文件做有限的修正操作.该参数的值即为允许的修正操作的名称.多个名称之间用英文半角逗号分隔
```
## 6. fmt与gofmt 命令
> Go 开发团队不想要 Go 语言像许多其它语言那样总是在为代码风格而引发无休止的争论,浪费大量宝贵的开发时间,因此他们制作了一个工具:go fmt（gofmt）
```sh
fmt与gofmt 命令 格式化go源文件,fmt命令实际"gofmt -l -w"命令之上做了一层包装,我们一般使用
usage: go fmt [-n] [-x] [packages]
 用法: go fmt [-n] [-x] 包
 可选参数:
    -x 打印执行的命令
    -n 打印执行的命令,但不真正执行

示例:
# 格式化 testmod 包, 并显示执行命令
go fmt -x testmod


gofmt 命令
usage: gofmt [flags] [path ...]
 用法: gofmt [参数] [路径 ...]
 可选参数:
    -cpuprofile string 将cpu配置文件写入此文件
    -d 显示格式化前后差异,但不写入文件
    -e 打印所有错误, 默认只会打印不同行的前10个错误
    -l 列出需要格式化的文件
    -r string 重新规则,方便我们做批量替换,例如我们需要把hellomod.Hello替换成hellomod.HelloNew("hellomod.Hello -> hellomod.HelloNew")
    -s 简化代码
    -w 将结果直接写入到文件中
    
示例:
# 格式当前目录代码
gofmt -w ./

# 把当前目录中的“hellomod.Hello” 替换成 "hellomod.HelloNew"
gofmt -r "hellomod.Hello -> hellomod.HelloNew" -w ./
```


## 8. get 命令 下载并安装包和依赖(下载包和依赖,并对它们进行编译安装)
```sh
get 命令 下载并安装包和依赖(下载包和依赖,并对它们进行编译安装)
usage: go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]
 用法: go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [包]
 可选参数:
    -d 只下载不安装(只执行下载动作, 不执行安装动作)
    -f 只有在包含了-u参数的时候才有效.该参数会让命令程序忽略掉对已下载代码包的导入路径的检查.如果下载并安装的代码包所属的项目是你从别人那里Fork过来的,那么这样做就尤为重要了
    -fix 会下载代码包后先执行修正动作,而后再进行编译和安装
    -insecure 请谨慎使用, 允许使用不安全(http或者自定义域)的存储库中下载解析.
        即:允许命令程序使用非安全的scheme（如HTTP）去下载指定的代码包.如果你用的代码仓库(如公司内部的Gitlab)没有HTTPS支持,可以添加此标记.请在确定安全的情况下使用它.
    -t 同时也下载需要为运行测试所需要的包
    -u 强制从网络更新包和它的依赖包.默认情况下,该命令只会从网络上下载本地不存在的代码包,而不会更新已有的代码包
    -v 显示执行的命令

示例:
# 下载包
go get github.com/donvito/hellomod
```

## 9. install 命令 编译并安装指定包及它们的依赖包,
```sh
install 编译并安装指定包及它们的依赖包,先生成中间文件(可执行文件或者.a包),然后把编译好的结果移到$GOPATH/pkg或者$GOPATH/bin
usage: go install [-i] [build flags] [packages]
 用法: go install [-i] [编译 flags] [包]
 可选参数:
    -i 

示例:
# 安装包
go install github.com/gin-gonic/gin
```

## 10. list 命令 列出包和模块信息
```sh
list 列出包和模块信息
usage: go list [-f format] [-json] [-m] [list flags] [build flags] [packages]
 用法: go list [-f format] [-json] [-m] [list flags] [build flags] [包]
 可选参数:
    -f {{.字段名}} 查看指定的字段信息
    -json 以json格式打印信息
    -m 列出模块信息
更多用法请参考(go help list)

示例:
# 以json格式打印gapp包信息
go list -json gapp

# 打印模块信息
go list -m testmod

# 以json格式打印模块信息
go list -m -json testmod 
# testmod模块打印结果:
{
        "Path": "testmod",
        "Main": true,
        "Dir": "/Users/zhaoweijie/go/src/business-card/docker-compose/go-tutorials/9/examples/testmod",
        "GoMod": "/Users/zhaoweijie/go/src/business-card/docker-compose/go-tutorials/9/examples/testmod/go.mod"
}

```


## 13. go test

```sh
go test
    -c 把包编译二进制测试包, 注意不会运行, 需要你手动执行二进制测试
        示例: 
            go test -c package_import_path
            go test -c 包的导入路径
            1. go test -c 在当前目录生成二进制测试
            2. go test -c go test -c go-tutorials/8/examples/demo1 

    -exec 运行二进制测试
        示例:
            go test -c -exec demo1.test
    -json 运行测试,并将结果输出为json格式
        示例:
            go test -json path
            1. go test -json 测试当前包
            2. go test -json ./
    -o 把测试编译成自己命名二进制包, 默认仍然会运行测试(除非指定-c或者-i)
        示例:
            go test -o file_name
            go test -o 文件名
            1. go test -o demo1.custom_name.test
    -bench 运行基准测试, 默认情况下不运行
        示例:
            go test -bench regexp
            go test -bench 正则表达式
            1. go test -bench 运行所有基准测试
            2. go test -bench=. 运行所有基准测试
            3. go test -bench=hah 运行指定的基准测试
    -benchtime 基准测试,持续时间(默认1秒)
    
    -count 运行测试次数
        示例:
            go test -count n
            go test -count 次数
            1. go test -count 10 运行所有的测试10次
    -cover 覆盖率统计, 注意覆盖率统计是通过代码注释来工作的
    -cpu 指定测试cpu数量
        示例:
            go test -cpu 1,2,4
            go test -cpu cpu数量
            1. go test -cpu 8 指定8个cpu
    -list regexp 列出匹配的测试
        示例:
            go test -list regexp
            go test -list 正则表达式
            1. go test -list Login 列出demo1中的测试
    -v 详细输出运行时记录所有的测试
        示例:
            go test -v
```