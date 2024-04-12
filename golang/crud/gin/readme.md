# 项目名称

## 项目简介

这是一个使用 Go 语言和 Gin 框架开发的 web 服务项目。项目主要包含以下功能：

- [x] 权限验证：使用 Bearer Token 进行权限验证。
- [x] 异常处理：统一处理应用程序的异常。
- [ ] 账户管理：包括账户的创建、查询、更新和删除。

## 目录结构

以下是项目的基本目录结构：

```
.
├── cmd
│   ├── server      # 服务启动目录
├── internal        # 内部代码
│   ├── api         # API 接口
│   ├── config      # 系统配置或加载
│   ├── constants   # 常量定义
│   ├── controller  # 控制器
│   ├── datasource  # 数据源
│   ├── error       # 错误处理
│   ├── middleware  # 中间件
│   ├── model       # 数据模型
│   ├── repository  # 数据仓库
│   ├── service     # 服务
├── pkg             # 通用包
├── sql             # 数据库脚本
├── web             # 静态资源
├── .env            # 环境变量
├── .example.env    # 环境变量示例
├── .gitignore      # Git 忽略文件
├── README.md       # 项目说明
├── main.go         # 项目入口
├── go.mod          # 项目依赖
└── go.mod          # 项目依赖
```

## 如何运行

在项目根目录下运行以下命令：

```bash
go run .
```

## 如何测试

在项目根目录下运行以下命令：

```bash
go test ./...
```
