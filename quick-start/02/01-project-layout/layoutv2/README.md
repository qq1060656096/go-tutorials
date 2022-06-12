### Go 工程化
```sh
1. 工程目录结构
2. API 设计
3. 配置管理
4. 测试
```
* **已定稿** [1. API设计](01/README.md)
* **已定稿** [2. 配置管理](02/README.md)
* **已定稿** [3. 包管理](02/README.md)
* **已定稿** [3. 测试](02/README.md)



# 安装protoc编译插件和grpc生成插件
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest



```
# 生成go pb 文件
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
# 生成go grpc service
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1



# 生成 go pb文件
protoc --proto_path=./ --go_out=./ --go_opt=paths=source_relative  *.proto

# 生成 php pb 文件
protoc --php_out=./ order.proto
protoc --php-grpc_out=./ ./order.proto

# 生成 pb 文件和 grpc 服务端
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative account/v1/account.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. ./account/v1/account.proto
# 生成 go pb 代码
protoc --go_out=./ ./account/account.proto
# 生成 grpc 服务端代码
protoc --go-grpc_out=. ./account/account.proto



# 生成 go pb 代码
protoc --go_out=./ account/v1/account.proto
# 生成 grpc.pb.go
protoc --go-grpc_out=. account/v1/account.proto




```sh
# 启动服务端
go run grpc-server.go 

# 启动客户端
go run grpc-client.go 
```
