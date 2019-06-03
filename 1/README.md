## go语言介绍

```
1. go语言介绍
2. go语言特点
3. go语言用途
4. go语言案例
5. go语言安装
6. go语言"hello world"
```
#### 1. go语言介绍
```
Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。
Go是从2007年末由Robert Griesemer, Rob Pike, Ken Thompson主持开发，后来还加入了Ian Lance Taylor, Russ Cox等人，并最终于2009年11月开源，在2012年早些时候发布了Go 1稳定版本。现在Go的开发已经是完全开放的，并且拥有一个活跃的社区。
```

#### 2. go语言特点
```
1. 语法简洁
2. 开发效率高
3. 安全
4. 并行、有趣
5. 内存管理
6. 数组安全
7. 编译迅速
8. 开源
```

### 3. go语言用途
```go
1. 服务开发
2. web(网页和api)开发
3. 分布式服务集群
```

### 4. go案例
```go
1. go
  go语言本身也是go语言写的
2. docker
    虚拟华平台
3. kubernetes
    用于调度和管理[Docker]
4. ngrok(内网穿透)
    是一个反向代理,它能够让你本地的web服务或tcp服务通过公共的端口和外部建立一个安全的通道,使得外网可以访问本地的计算机服务.
5. lantern
    蓝灯,一款P2P的翻墙软件
6. shadowsocks-go
    翻墙工具
```
### 5. go语言安装
> go语言安装不做过来介绍，请自行搜索或者点击链接
[go语言安装教程](https://www.runoob.com/go/go-environment.html)

### 6. go语言"hello world"
> 1. 创建文件hello.go, 并写入以下内容
```
package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}
```
> 2. 执行: go run hello.go
![1611558776058_.pic_hd.jpg](https://upload-images.jianshu.io/upload_images/6713312-31c2d7fd819c2e9b.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
