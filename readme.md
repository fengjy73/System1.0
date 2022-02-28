# 第一阶段
```txt
主要是User Service
```

## API
```txt
注意到只有管理员可以创建账户，用户注册实际是调用超级管理员创建账户
并且一旦用户创建，username不再改变
```
### Login\Logout
登入，利用cookie

1.1 POST/api/auth/login

登出

1.2 POST/api/auth/logout

返回当前用户

1.3 GET/api/auth/whoami

2.User

创建用户

2.1 POST /api/user/create

获取特定用户信息

2.2 GET /api/user

获取用户列表

2.3 GET /api/user/list

修改根据用户名特定用户信息

2.4 POST /api/user/update

软删除一个用户

2.5 POST /api/user/delete

## For Backend Specially

Basic requirement is as below.
```
Language: Golang
Database: Mysql
    127.0.0.1
    root
    123456
Connecting DB: Gorm
Sending Messages to FrontEnd: Gin
Structure:
1.Types(Given)
2.Dao
3.Service
4.Controller
5.Router 
```
详情见Types.go
