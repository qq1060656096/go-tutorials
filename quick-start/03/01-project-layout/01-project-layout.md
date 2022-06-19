

> 我们尝试学习 Go，或者我们正在为自己建立一个 Go 或一个玩具项目，这个项目布局是没啥必要的。从一些非常简单的事情开始(一个 main.go 文件绰绰有余)。当有更多的人参与这个项目时，你将需要更多的结构，包括需要一个 toolkit 来方便生成项目的模板，尽可能大家统一的工程目录布局

```shell
layoutv2          根目录
├─api               放置 API 的定义(protobuf)，以及对应的生成的 client 代码，基于 pb 生成的 swagger.json
│  ├─google             公司名
│     ├─accounts            服务名（这里演示服务名：账户服务）
│       ├─v1                    服务v1版本
│       ├─v2                    服务v2版本
├─cmd               命令存放目录本质负责启动、关闭、配置、初始化等
│  ├─account-interface  对外的 BFF 服务，接受来自用户的请求，比如暴露了 HTTP/gRPC 接口
│  ├─accounts-admin     admin 更多是面向运营测的服务，通常数据权限更高，隔离带来更好的代码级别安全。
│  ├─accounts-job       常驻任务，消费消息中间件
│  ├─accounts-service   service 对内的微服务，仅接受来自内部其他服务或者网关的请求，比如暴露了gRPC 接口只对内服务。
│  ├─accounts-task      定时任务，类似 cronjob
├─configs           配置文件模板或默认配置，比如database.yaml、redis.yaml、application.yaml。，比如database.yaml、redis.yaml、application.yaml 等
├─internal          应用私有库代码（您不希望其他人在其应用程序或库中导入的代码。请注意，这种布局模式是由 Go 编译器本身强制执行的。）
│  ├─accounts           账号服务目录（/internal/app 我们们习惯把相关的服务放在单独的一个目录）
│     ├─biz             业务逻辑目录，类似 DDD 的 domain
│     ├─data            数据访问目录，类似 DDD 的 repository 仓储
│     ├─pkg             服务内部可以使用的库代码
│     ├─server          存放 grpc、http 的一些代码
│     ├─service         服务目录实现了 API 的定义，类似 DDD Application
├─pkg               外部应用程序可以使用的库代码（/pkg/mypubliclib）
├─third_party       外部辅助工具，fork的代码和其他第三方工具（例如Swagger UI）
├─test              外部测试应用程序和测试数据。随时根据需要构建/test目录。对于较大的项目，有一个数据子目录更好一些。例如，如果需要Go忽略目录中的内容，则可以使用/test/data或/test/testdata这样的目录名字。请注意，Go还将忽略以“.”或“_”开头的目录或文件，因此可以更具灵活性的来命名测试数据目录。
├─go.mod         	Go 模块信息文件
├─Makefile         	makefile文件
└─README.txt        README文件
```


```sh
每个公司都应当为不同的微服务建立一个统一的 kit 工具包项目(基础库/框架) 和 app 项目。
基础库 kit 为独立项目，公司级建议只有一个，按照功能目录来拆分会带来不少的管理工作，因此建议合并整合。
```
kit 项目必须具备的特点:
统一
标准库方式布局
高度抽象
支持插件
