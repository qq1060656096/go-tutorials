## 常用命令

```sh
1. go run
2. go build
3. go test
```

###### 3. go test

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