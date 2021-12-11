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
1.{ "msg": "success"}
2. {"msg": "the id already exists"}


```

登录（post）

http://175.24.166.140:8080/login

```
服务端
{
    "id":6656666,
    "username":"1341233",
    "password":"626676",
    "mail":"4343431"
}

服务器
1.{"msg": "login successfully"}
2.{"msg": "the id or password is not true"}
```

忘记密码（put）

http://175.24.166.140:8080/forget

```
客户端
{
    "id":66288585866,
    "password":"6266857",
    "mail":"434853431"
}
服务器
1.{"msg": "success"}
2.{"msg": "the id or mail is not true"}
```

提问（post）

http://175.24.166.140:8080/hand_que

```
客户端
{
    "que":66288585866,
    "username":"6266857"
}
服务器
{"msg": "success"}
```

回答（post）

http://175.24.166.140:8080/hand_ans

```
客户端
{
    "ans":66288585866,
    "username":"6266857",
    "questionid":1
}
服务器
{"msg": "success"}
```

热搜（get）

http://175.24.166.140:8080/hot

```
客户端
{
    "great":66288585866,
    "userparam":"6266857",
    "que":1,
    "ansnum":1
}
服务器
{"hot":"" }
```

搜索（get）

http://175.24.166.140:8080/search

```
客户端
{
    "great":66288585866,
    "userparam":"6266857",
    "que":1,
    "ansnum":1
}
服务器
{"search": questions }
```

点赞（post）

http://175.24.166.140:8080/like

```
客户端
{
    "num":1,
   
}
服务器
客户端
{
    "great":66288585866,
    "userparam":"6266857",
    "que":1,
    "ansnum":1
}
服务器
{"msg": "ERROR" }
```