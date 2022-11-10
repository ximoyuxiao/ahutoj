# AHUTOJ代码提交规范

### 获取仓库代码

```sh
git clone + 仓库的链接
如：
git clone git@github.com:199094212/ahutoj.git
```

### 修改代码

在修改代码前 要提交的分支下 建立一个新的分支，在新的分支改完代码后才能提交

新分支的命名方式：学号/原分支名/修改内容

如我要在master分支创建一个login接口，我的做法应该是

```sh
git checkout master # 切换会原来的master分治
git pull   #拉取master分支在远程的代码
git checkout -b 199094212/master/add-login #在现在分支的基础上，创建新的分支
git push origin 199094212/master/add-login # 在远程仓库建立这个分支
git branch --set-upstream-to=origin/199094212/master/add-login 199094212/master/add-login #本地分支远程分支做关联
```

### 提交代码

理论上要求每次代码的合并请求上只允许有一次提交

#### 代码提交

```sh
git add (文件1 文件2) || git add .(添加所有文件)
git commit -m '修改了什么'
git push
```

#### 如果要重新提交，但是提交信息不变

```sh
git add ... # 将新修改的文件 添加上去
git commit --amend # 将新的修改合并到上一次修改 后面会退出一个编辑框 那是让你修改提交信息的一般就 冒号 然后wq（vim）
```

#### 如果自己的分支落后了远程的主分支

```sh
git checkout master
git pull
git checkout 199094212/master/add-login 
git push
```

### 接口设计

### 接口文档

接口文档要求

接口地址：/api/.....

接口的请求方法GET|POST|PUT|DELETE

接口的数据格式：一般是：x-www-form-urlencoded

接口的请求参数（以json形式表示）

```json
{
    username:"admin"  #UID
    password:"199094212" #密码
}
```

接口的返回参数

```json
错误状态的code和msg单独列出如：
code 1001 msg:"参数错误"
此处罗列出返回的对象的信息（可以用postman返回后 直接拷贝过来）如下
{
    code: 0 
    msg: "sucess"
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYWRtaW4iLCJleHAiOjE2NTcxOTcyNzQsImlzcyI6ImFodXRvaiJ9.noGXyFe8JE0Pd6wOfOS71NyuAFry0BWVskaL_H9mlfg"
}
```



