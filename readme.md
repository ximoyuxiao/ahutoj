---
typora-root-url: img
---

# AHUT在线判题系统
### 环境:
go 1.18.3

ubuntu18.06
### 在本机中跑起这个项目
1、执行install.sh (其会自动下载Mysql、redis)，自动执行sql脚本建立数据库表）

2、配置config.yaml文件（具体可以参考config.yaml.bak）

3、执行make run 既可以跑起来项目

### 相关依赖:
    C/C++:rabbitMQ\libamqp-dev,hiredis,libmysql
    Go:gorm、gin、air

### 项目结构图
    .
    ├── doc             项目文档
    |—— core            判题模块
    |—— gatway          网关服务
    |—— originJudge     外部判题服务
    |—— originproblem   爬取第三方题目
    |—— useranalytics   用户分析
    └── web
        ├── dao  数据库层
        ├── io   IO层（请求和响应）
        │   ├── constanct  一些常量
        │   ├── request    请求
        │   └── response   响应
        ├── logic          逻辑层 用于处理业务逻辑
        ├── middlewares    中间件层
        ├── models         模型层，主要是一些视图
        ├── routers        路由层，决定走哪个服务
        ├── service        服务层
        └── utils          一些常用工具
### 项目分层模型
    =====================
    |网关层             |
    ====================
    |服务层             |
    ====================
    |逻辑层             |
    ====================
    |模型层             |
    ====================
    |数据库层           |
    ====================
### 实现模块

#### 用户模块
    1.用户注册模块 register
    2.用户登录模块 login
    3.用户注销模块 exit
    4.用户管理模块 info
    5.绑定Vjudge账户
#### 管理员模块

    1.用户权限管理模块 
    2.题目管理模块
    3.竞赛管理模块
    4.获取用户做题情况模块
#### 题目模块
    1.题目列表模块
    2.题目展示模块
    3.用户题单模块
#### 竞赛系统模块
    1.添加竞赛模块
    2.删除竞赛模块
    3.获取竞赛信息模块
    4.修改竞赛模块
#### 判题系统模块
    1.内部判题模块     judged  (C++实现）
    2.第三方判题模块   originJudge(go实现)
## 三、数据库及其架构设计

数据库设计

<img src="./img/sql.png" alt="sql" style="zoom: 80%;" />
[查看数据库代码：](./doc/oj.sql)