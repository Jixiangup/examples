# GIN Restful API Example

这是使用 GIN 框架的 RESTful API 的简单示例。

## 项目结构

```
.
├── README.md       # 项目说明文档
├── go.mod          # Go modules 的依赖管理文件
├── go.sum          # Go modules 的依赖校验文件
├── main.go         # 项目的入口文件
├── cmd             # 应用程序的主要执行文件
│   └── server
│       └── main.go # 服务器的主要执行文件，负责启动服务器
├── internal        # internal 目录包含了项目的内部包，这些包只能被当前项目使用
│   ├── api
│   │   └── router.go # 路由配置文件，定义了所有 API 路由和对应的处理函数
│   └── config
│       └── config.go # 配置文件，包含了项目运行所需的配置信息，如数据库连接信息
└── pkg             # pkg 目录包含了可以被其他项目使用的库代码
    └── error       # 错误处理相关的代码，提供了统一的错误处理机制
```
