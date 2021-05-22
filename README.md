# go-websocket
基于golang实现的分布式聊天系统，支持一对一聊天，聊天室等功能。为了测试方便发送消息数据暂未存入数据库，后期会加入数据库，也可自行加入数据库，方便永久存储聊天内容，以及支持消息必达等功能。

### 依赖包

```
github.com/go-redis/redis
github.com/gin-gonic/gin
github.com/gorilla/websocket
github.com/smallnest/rpcx
```

包说明：

  redis ：用于缓存ws服务器信息，用心跳形式维护ws服务器信息。

 gin：实现web服务

 websocket：  实现websocket协议

rpcx：服务器建rpc通信

### 架构图

![](https://github.com/guyan0319/go-websocket/blob/main/doc/jiagoutu.png?raw=true)

### 一对一发消息

- 客户端 发送建立长连接请求，经过nginx负载均衡分配给其中一台ws服务器（这里假设分配的是A服务器）处理。

- A服务器响应长连接请求，并缓存客户端地址和用户连接，用户id等信息。

- 客户端收到服务端响应建立websocet连接。

- 客户端发送信息，nginx负载均衡分配给其中一台ws服务器（这里假设是B服务器）。

- B服务器从发送的信息中解析接收用户（假设为a）信息，先验证a用户是否和B服务器建立websocet连接，若建立则直接发送消息给a用户。否则通过redis缓存中获取ws服务器信息列表，通过rpc方式发送消息到ws服务器列表中除B服务器之外的每台ws服务器，这些接收到发送信息的ws服务器，先验证和a用户是否建立连接，建立则发送信息给a用户，否则丢弃。

### 群发消息

  - 客户端 发送建立长连接请求，经过nginx负载均衡分配给其中一台ws服务器（这里假设分配的是A服务器）处理。
  - A服务器响应长连接请求，并缓存客户端地址和用户连接，用户id等信息。
  - 客户端收到服务端响应建立websocet连接。
  - 客户端发送信息，nginx负载均衡分配给其中一台ws服务器（这里假设是B服务器）。
  - B服务器从发送的信息中解析出群信息，根据群信息获取用户列表，遍历用户发送信息（发送方式跟一对一类似）。
  - 先验证用户是否和B服务器建立websocet连接，若建立则直接发送消息给用户。否则通过redis缓存中获取ws服务器信息列表，通过rpc方式发送消息到ws服务器列表中除B服务器之外的每台ws服务器，这些接收到发送信息的ws服务器，先验证和用户是否建立连接，建立则发送信息给用户，否则丢弃。

###   快速搭建

  1、拉取代码

```
git clone https://github.com/guyan0319/go-websocket.git
```

注：这里代码版本控制使用[go modules](https://github.com/guyan0319/golang_development_notes/blob/master/zh/1.10.md) 

2、运行系统

```
go run main.go
```

3、配置nginx

```
upstream  go-http
{
    server 127.0.0.1:8282 weight=1 max_fails=2 fail_timeout=10s;
    keepalive 16;
}

upstream  go-ws
{
    server 127.0.0.1:8089 weight=1 max_fails=2 fail_timeout=10s;
    keepalive 16;
}

server {
        listen        80;
        server_name  ws.test;
        root   "D:/phpstudy/WWW/";
          location /ws {
        proxy_set_header Host $host;
        proxy_pass http://go-ws;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection $connection_upgrade;
        proxy_set_header Connection "";
        proxy_redirect off;
        proxy_intercept_errors on;
        client_max_body_size 10m;
    }

    location /
    {
	    proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host $host;
        proxy_pass http://go-http;
    }
}

```

4、测试一对一聊天

浏览器打开两个窗口访问http://ws.test/home/index?uid=1&to_uid=2

http://ws.test/home/index?uid=2&to_uid=1

5、测试群聊天





相关资料：

github.com/gorilla/websocket

https://my.oschina.net/u/4231722/blog/3168223

https://doc.rpcx.io/

https://github.com/link1st/gowebsocket

https://segmentfault.com/a/1190000018712908

