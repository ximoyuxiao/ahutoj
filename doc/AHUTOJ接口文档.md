## AHUTOJ接口文档

#### 用户登录接口
接口地址：/api/auth/login/
接口的请求方法:POST
接口的数据格式：json
接口的请求参数:
```json
{
    uid: string
    pass: string
}
```
接口的返回值
```json
{
    "code": 0,
    "msg": "success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYWRtaW4iLCJleHAiOjE2NTczNzYxMzEsImlzcyI6ImFodXRvaiJ9.DnM0dQ0BDDcH78PUV50PxfQwg7dLAaQovOcpvMeTaO0"
}
```
#### 用户注册接口
接口地址：/api/auth/register/
接口的请求方法:POST
接口的数据格式：json
接口的请求参数:
```json
{
    uid: string  用户ID
    uname:string 用户名称
    pass:string  用户密码
    school:string
    classes:string
    major:string
    vjid:string
    vjpwd:string
}
```
