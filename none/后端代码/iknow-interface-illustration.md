# iknow 接口文档

注册（post）

http://175.24.166.140:8080/register

```
客户端
{
    "id":6656666,
    "username":"1341233",
    "password":"626676",
    "mail":"4343431"
}
服务器
1.{ "msg": "register successfully"}
2. {"msg": "the id already exists"}


```

登录（post）

http://175.24.166.140:8080/login

```
客户端
{
    "id":6656666,
    "username":"1341233",
    "password":"626676",
    "mail":"4343431"
}

服务器
1.
{
    "msg": "login successfully"
}{
    "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjAwMDMsImV4cCI6MTYzOTI4NDI4NCwiaXNzIjoiaWtub3ciLCJuYmYiOjE2MzkyODQwNDR9.GUPbDJqJAk6byyxjT0QEs_vLj-uHeAR_8L0rPBwuafE",
    "msg": "创建token成功"
}



2.
{
    "msg": "the id or password is not true"
}{
    "msg": "鉴权失败"
}



3.
{
    "msg": "创建token失败"
}
```

登录--jwt--token认证（get）

http://175.24.166.140:8080/home 

```
客户端  --Authorization--type：Bearer Token
Token   （过期时间为3minute）
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjAwMywiZXhwIjoxNjM5MjgzODcyLCJpc3MiOiJpa25vdyIsIm5iZiI6MTYzOTI4MzgwN30.YsansvJn9ltniSw1gN9WhGem5sTEhmMuWVrBLJIVnVw
服务器
1.
{
"msg":  "verify successfully",
		"id":id,
}

2.
{
    "code": 2005,
    "msg": "无效的Token"
}

3.
{
    "code": 2003,
    "msg": "请求头中auth为空"
}

4.
{
    "code": 2004,
    "msg": "请求头中auth格式有误"
}
```

删除ID（put）

http://175.24.166.140:8080/del_id

```
客户端
{
    "id":2003    
}
服务器
1.Headers中的code错误
{
    "message": "无权访问（密码是产品名字）"
}
2.Headers中的code正确
{
    "msg": "delete successfully"
}

```

忘记密码（put）

http://175.24.166.140:8080/forget

```
客户端
{
    "id":88998,
    "newpassword":"616",
    "mail":"666"
}
服务器
1.
{
    "msg": "reset password successfully"
}

2.
{
    "msg": "the id or mail is not true"
}
```

提问（post）

http://175.24.166.140:8080/hand_que

```
客户端
{
    "que":"66288585866",
    "username":"6266857"
}
服务器
{
    "msg": "handle question successfully"
}
```

删除问题（post）

http://175.24.166.140:8080/del_que

```
客户端
{
   "id":5
}
服务器
1.Headers中的code错误
{
    "message": "无权访问（密码是产品名字）"
}
2.Headers中的code正确
{
    "msg": "delete question successfully"
}
```

回答（post）

http://175.24.166.140:8080/hand_ans

```
客户端
{
    "ans":"66288585866",
    "username":"6266857",
    "questionid":1
}
服务器
1.Headers中的code错误
{
    "message": "无权访问（密码是产品名字）"
}
2.Headers中的code正确
{
    "msg": "hand answer successfully"
}
```

删除答案（post）

http://175.24.166.140:8080/del_ans

```
客户端
{
    "id":5
}
服务器
1.Headers中的code错误
{
    "message": "无权访问（密码是产品名字）"
}
2.Headers中的code正确
{
    "msg": " delete answer successfully"
}
```

热搜（get）

http://175.24.166.140:8080/hot

```
客户端


服务器
{"hot":show }
```

搜索（get）

http://175.24.166.140:8080/search

```
客户端
{
    "keywords":"66288585866",
}
服务器
1.{"msg":"error"}
2.{"search":upload}
```

点赞（post）

http://175.24.166.140:8080/like

```
客户端
{
    "questionid":1,
   
}
服务器
1.{"msg":"error"}
2.{"msg":"点赞成功"}
```

收藏（post）

http://175.24.166.140:8080/collect

```
客户端
{
    "userid":1,
    "questinid":1
   
}
服务器
1.{"msg":"error"}
2.{"msg": "collect successfully"}
```

