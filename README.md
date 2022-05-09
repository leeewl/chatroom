# 一个聊天室的项目

## 项目介绍
这是一款基于websocket的聊天软件，后端需要部署到服务器，使用golang + postgresql开发。
前端用浏览器打开，从`http://域名或ip:端口/login`进入。域名，端口要修改config.yml文件。


## 用户使用说明

* 用户可以在login界面注册账号，然后登录
* 登陆后会跳转到聊天界面
* 线进入房间，再发送聊天信息

## 界面
聊天界面介绍
![](https://github.com/leeewl/chatroom/blob/master/images/introduce.png)
现在重点是后端，前端走极丑风
![](https://github.com/leeewl/chatroom/blob/master/images/interface_chat.png)


## 进度

玩家可以进入不同的聊天室聊天,可以显示在线玩家。

## 部署

### 数据库

#### postgresql

1. 创建数据库用户名和密码，需要分别与 chatroom/conf/systemConf.go里面的常量`DB_USER`,`DB_PASSWORD`对应。

2. 创建数据库

```
create database chatroom owner gopher;
```

3. 创建数据库表

```
 psql -U YOUR_USER_NAME -d chatroom -f 项目所在目录/chatroom/module/room/sql/t_room.sql
 psql -U YOUR_USER_NAME -d chatroom -f 项目所在目录/chatroom/module/user/sql/t_user.sql

```

