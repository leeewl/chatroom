# 做一个聊天室的项目

## 玩家登录的流程

先登录界面，登陆后进入聊天界面，然后选择房间

## 界面
现在重点是后端，前端走极丑风
![](https://github.com/leeewl/chatroom/blob/master/images/interface_chat.png)

## 进度

玩家可以进入不同的聊天室聊天

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

