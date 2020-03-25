## go + c混合开发

```
1. 直接嵌套在go文件中使用,最简单直观的
2. 导入动态库 .so 或 dll 的形式, 最安全但是很不爽也比较慢的
3. 直接引用 c/c++ 文件的形式,层次分明,容易随时修改看结果的
```

### 1. 直接嵌套在go文件中使用,最简单直观的
```
1. 但凡要引用与 c/c++ 相关的内容，写到 go 文件的头部注释里面
2. 嵌套的 c/c++ 代码必须符合其语法，不与 go 一样
3. import "C" 这句话要紧随，注释后，不要换行，否则报错
4. go 代码中调用 c/c++ 的格式是: C.xxx()，例如 C.add(2, 1)
```

```
go run examples/demo1/main.go
go.call.Add(1, 2)=3 
c hello !%  
```


### 2. 导入动态库 .so 或 dll 的形式, 最安全但是很不爽也比较慢的
```
1.编译动态链接库
gcc -c examples/demo2/lib/demo_lib.c -o examples/demo2/lib/demo_lib.o

gcc -c -fpic examples/demo2/lib/demo_lib.c -o examples/demo2/lib/demo_lib.o

gcc -shared examples/demo2/lib/demo_lib.o -o examples/demo2/lib/demo_lib.so

gcc examples/demo2/lib/demo.c -fPIC -shared -o examples/demo2/lib/demo.so

gcc -shared -o examples/demo2/lib/demo_lib.so examples/demo2/lib/demo_lib.o

2. go中通过注释指明动态链接库路径
3. go中使用动态链接库
go run examples/demo2/main.go
```

```
先回答为什么说这种是最安全的和最不爽的？原因如下:
  1. 动态库破解十分困难, 如果你的 go 代码泄露，核心动态库没那么容易被攻破
  2. 动态库会在被使用的时候被加载, 影响速度
  3. 操作难度比方式一麻烦不少
结论
CFLAGS: -I路径 这句话指明头文件所在路径, -Iinclude 指明 当前项目根目录的 include 文件夹
LDFLAGS: -L路径 -l名字 指明动态库的所在路径，-Llib -llibvideo, 指明在 lib 下面以及它的名字 video
如果动态库不存在,将会爆找不到定义之类的错误信息
```



### 3. 直接引用 c/c++ 文件的形式,层次分明,容易随时修改看结果的
```
go run examples/demo3/main.go
go.call.c.file.demo
Add(3, 30)=33 
```

[参考文章- 全面总结： Golang 调用 C/C++，例子式教程](https://blog.csdn.net/zdy0_2004/article/details/79124269)