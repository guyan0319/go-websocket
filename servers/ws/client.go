package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

// Client is a websocket client
type Client struct {
	Addr          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接
	Send          chan []byte     // 待发送的数据
	AppId         uint32          // 登录的平台Id
	UserId        string          // 用户Id
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间
}

// Message is return msg
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

func (c *Client) Read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("read stop", string(debug.Stack()), r)
		}
	}()
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		c.Socket.PongHandler()
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		//处理客户端信息
		log.Printf("读取到客户端的信息:%s", string(message))
		Handle(c,message)
	}
}

func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Printf("发送到到客户端的信息:%s", string(message))
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

//TestHandler socket 连接 中间件 作用:升级协议,用户验证,自定义信息等
func WsPage(c *gin.Context) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	conn.CloseHandler()
	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())
	currentTime := uint64(time.Now().Unix())
	//可以添加用户信息验证
	client := &Client{
		Addr:          conn.RemoteAddr().String(),
		Socket:        conn,
		Send:          make(chan []byte, 100),
		FirstTime:     currentTime,
		HeartbeatTime: currentTime,
	}
	go client.Read()
	go client.Write()
	Manager.Register <- client
}

// 读取客户端数据
func (c *Client) SendMsg(msg []byte) {

	if c == nil {

		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()

	c.Send <- msg
}

// 读取客户端数据
func (c *Client) close() {
	close(c.Send)
}

// 用户登录
func (c *Client) Login(appId uint32, userId string, loginTime uint64) {
	c.AppId = appId
	c.UserId = userId
	c.LoginTime = loginTime
	// 登录成功=心跳一次
	c.Heartbeat(loginTime)
}

// 用户心跳
func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime

	return
}

// 心跳超时
func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}

	return
}

// 是否登录了
func (c *Client) IsLogin() (isLogin bool) {

	// 用户登录了
	if c.UserId != "" {
		isLogin = true

		return
	}

	return
}

// 读取客户端数据
func (c *Client) GetKey() (key string) {
	key = GetUserKey(c.AppId, c.UserId)

	return
}


