---
typora-root-url: img
---

# AHUT在线判题系统
### 环境:
go 1.18.3
ubuntu18.06
### 在本机中跑起这个项目
1、执行install.sh (其会自动下载Mysql、redis)，自动执行sql脚本建立数据库表）（开发中。。。）
2、配置config.yaml文件（具体可以参考config.yaml.bak）
3、执行make run 既可以跑起来项目

### 项目结构图
    .
    ├── doc   项目文档
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
    1.代码提交模块 submit
    2.判题模块     judged  (C++实现）

## 三、数据库及其架构设计

数据库设计

<img src="./img/sql.png" alt="sql" style="zoom: 80%;" />

​	                                                                                                           图1

### 1.UserTable 用于记录用户的基本信息

| 意义   | 变量    | 类型         | 条件        |
| ------ | :------ | :----------- | :---------- |
| 学号   | uid     | varchar(20)  | primary key |
| 姓名   | name    | varchar(20)  |             |
| 密码   | pass    | varchar(128) |             |
| 学校   | school  | varchar(128) |             |
| 班级   | classes | varchar(20)  |             |
| 专业   | major   | varchar(30)  |             |
| vj账号 | vjid    | varchar(30)  |             |
| vj密码 | vjpwd   | varchar(128) |             |
| 邮箱   | E_mail  | varchar(20)  |             |
____
### 2.PermissionTable 用户权限列表(包含管理员权限，举办比赛权限，查看代码权限，用户管理权限)
| 意义         | 变量            | 类型        | 条件                    |
| ------------ | :-------------- | :---------- | :---------------------- |
| 学号         | uid             | varchar(20) | reference userTable(id) |
| 管理员权限   | administrator   | varchar(2)  | in ('N','Y')            |
| 题目编辑权限 | problem_edit    | varchar(2)  | in ('N','Y')            |
| 阅读源码权限 | source_browser  | varchar(2)  | in ('N','Y')            |
| 创建比赛权限 | contest_creator | varchar(2)  | in ('N','Y')            |
___
### 3.ProblemTable 问题列表3
ProblemTable(<u>问题id</u> 标题，描述，输入，输出，样例输入，样例输出，提示，极限时间，极限内存)

| 意义     | 变量          | 类型         | 条件                        |
| -------- | :------------ | :----------- | :-------------------------- |
| 题目编号 | pid           | int          | primary_key  AUTO_INCREMENT |
| 标题     | title         | varchar(128) |                             |
| 描述     | description   | TEXT         |                             |
| 输入     | input         | TEXT         |                             |
| 输出     | output        | TEXT         |                             |
| 样例输入 | sample_input  | TEXT         |                             |
| 样例输出 | sample_output | TEXT         |                             |
| 提示     | hit           | TEXT         |                             |
| 极限时间 | limitTime     | int          |                             |
| 极限内存 | limitMemory   | int          |                             |
___
### 4.ListTable    题单列表
| 意义       | 变量  | 类型        | 条件                        |
| ---------- | :---- | :---------- | :-------------------------- |
| 题单号     | lid   | int         | primary_key  AUTO_INCREMENT |
| 创建者学号 | uid   | varchar(20) | reference UserTable(uid)    |
| 标题       | title | Text        |                             |
| 创建时间   | stime | datetime    |                             |

### 5.ListProblemTable 题单和题目 n：m关系
| 意义     | 变量 | 类型 | 条件                        |
| -------- | :--- | :--- | :-------------------------- |
| 题单编号 | lid  | int  | reference ListTable(lid)    |
| 题目编号 | pid  | int  | reference ProblemTable(pid) |
___

### 6.ListUserTable 题单用户表 m：n关系
ListUserTable(题单编号，学号，提交数量，ac数量)
| 意义     | 变量       | 类型        | 条件                     |
| -------- | :--------- | :---------- | :----------------------- |
| 题单编号 | lid        | int         | reference ListTable(lid) |
| 用户编号 | uid        | varchar(20) | reference UserTable(uid) |
| 提交次数 | submit_num | int         |                          |
| 通过题数 | ac_num     | int         |                          |

### 7.ContestTable 竞赛列表

| 意义                     | 变量        | 类型         | 条件                        |
| ------------------------ | :---------- | :----------- | :-------------------------- |
| 竞赛编号                 | cid         | int          | primary_key  AUTO_INCREMENT |
| 学号                     | uid         | varchar(20)  | references UserTAble(uid)   |
| 竞赛标题                 | title       | Text         |                             |
| 竞赛描述                 | description | Text         |                             |
| 开始时间                 | begin_time  | datetime     |                             |
| 结束时间                 | end_time    | datetime     |                             |
| 竞赛类型                 | ctype       | varchar(15)  | in('ACM','OI')              |
| 竞赛公开                 | ispublic    | varchar(10)  | in(private,public)          |
| 参赛密码[私有情况下存在] | pass        | varchar(128) |                             |
___

### 8.ConProTable竞赛问题列表 n:m
| 意义     | 变量       | 类型 | 条件                           |
| -------- | :--------- | :--- | :----------------------------- |
| 竞赛标号 | cid        | int  | references conpeleteTable(cid) |
| 题目编号 | pid        | int  | reference ProblemTable(pid)    |
| 提交数量 | submit_num | int  |                                |
| AC数量   | ac_num     | int  |                                |
### 9.SubmitTable  提交列表
| 意义                     | 变量       | 类型        | 条件                               |
| ------------------------ | :--------- | :---------- | :--------------------------------- |
| 提交编号                 | sid        | int         | primary_key  AUTO_INCREMENT        |
| 题目编号                 | pid        | int         | references problemList(pid)        |
| 学号                     | uid        | varchar(20) | references UserTable(uid)          |
| 竞赛编号[null表示非竞赛] | cid        | int         | references competeTable(cid)       |
| 判题机编号               | judgeid    | int         |                                    |
| 代码                     | source     | TEXT        |                                    |
| 语言                     | lang       | varchar(10) | in('C++11','JAVA','Python3','C99') |
| 结果                     | result     | varchar(10) |                                    |
| 花费时间                 | usetime    | int         |                                    |
| 花费内存                 | memory     | int         |                                    |
| 提交时间                 | submitTime | datetime    |                                    |
___

[查看数据库代码：](./doc/oj.sql)