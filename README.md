# MyEntry

> Shopee 入职项目 entry task

### 技术栈
* 前端框架 Layui 
* 语言: go
* 网络库: 
  * Http Server 标准库: net/http
  * Tcp Server 标准库: net/Tcp
* 日志: logrus : https://github.com/sirupsen/logrus
* mysql：https://github.com/go-sql-driver/mysql
* redis：https://github.com/go-redis/redis
* 包管理工具 [Go Modules](https://github.com/golang/go/wiki/Modules)
* 后台登录：
  * cookie 管理SeesionId
  * redis 管理SessionInfo
* 使用 YAML 文件进行多环境配置


### 目录结构

```shell
├── httpServer                   # http 服务根目录
│   ├── common                   # 初始化http服务
│   ├── conf                     # http 配置
│   ├── controller               # http 接口
│   ├── entity                   # 接口传输实体
│   ├── routers                  # 业务路由
│   ├── service                  # 接口服务层
│   ├── view                     # 模板引擎
│   └── main.go                  # http Server 项目入口文件
├── log                          # 存放日志的目录
├── pkg                          # 公共工具包
│   ├── content                  # 公共常量定义
│   ├── entity                   # 公共传输实体
│   ├── log                      # 日志工具包
│   ├── mysql                    # MySQL工具包
│   ├── pool                     # 通用连接池包
│   ├── redis                    # redis 工具包
│   ├── rpc                      # rpc C/S 基类
│   └── utils                    # 公共工具类
├── static                       # 静态文件
│   ├── js                       # js文件
│   ├── layui                    # layui框架文件
│   └── pic                      # 图片文件夹
├── tcpServer                    # tcp 服务根目录
│   ├── common                   # 初始化tcp服务
│   ├── mapper                   # DB服务层
│   ├── model                    # DB table 映射模型
│   ├── service                  # tcp 逻辑服务层
└───└── main.go                  # tcp Server 项目入口文件
```
### 部署流程

* 依赖环境：
    ####go  redis  mysql
       版本:   go version go1.17.3 darwin/amd64
              redis 4.0以上 或 brew install redis 默认安装最新版
              mysql 5.7及以上



* 安装部署 最新版即可
    ####redis 
        brew install redis
    ####mysql
        brew intall mysql
    ####项目部署
```shell
# 下载安装，可以不用是 GOPATH
git clone git@github.com:ShellBotCheng/my_entry.git

# 进入到下载目录
cd my_entry

# 修改环境配置文件
vim httpServer/conf/dev.yml   # http相关配置
vim tcpServer/conf/dev.yml    # tcp相关配置
# 修改 mysql、redis 配置

# 导入初始化 sql 结构
mysql -u{user} -p{pwd}
> create database entry_task;
> set names utf8mb4;
> use entry_task;
> source entry_task.sql;

# 生成测试数据 要求python 版本python 3及以上
    python my_entry.py    
# 下载依赖
    go mod tidy
# tcp 启动 
    go run ./tcpServer/main.go
# 或 tcp 后台启动
    nohup go run ./tcpServer/main.go &
# http 启动
    go run ./httpServer/main.go
# 或 http 后台启动
    nohup go run ./tcpServer/main.go &
```

* 访问
  http://localhost:8082
  
* 用户名：admin

* 密码：123456
