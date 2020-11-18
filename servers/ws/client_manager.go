package ws

import (
	"fmt"
	"go-websocket/lib/cache"
	"go-websocket/servers/msgs"
	"sync"
	"time"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients     map[*Client]bool   // 全部的连接
	ClientsLock sync.RWMutex       // 读写锁
	UserClients    map[string]*Client
	UserClientsLock    sync.RWMutex       // 读写锁
	SendUserMsg  chan *msgs.SendUserMsg //广播消息
	Register   chan *Client  //连接
	Login       chan *login        // 用户登录处理
	Unregister chan *Client //退出
}
// Manager define a ws server manager
var Manager = ClientManager{
	Clients:    make(map[*Client]bool),
	UserClients:      make(map[string]*Client),
	SendUserMsg:  make(chan *msgs.SendUserMsg),
	Register:   make(chan *Client),
	Login:      make(chan *login),
	Unregister: make(chan *Client),
}

// 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	manager.Clients[client] = true
}

// 删除客户端
func (manager *ClientManager) DelClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	delete(manager.Clients, client)
}

// 获取用户的连接
func (manager *ClientManager) GetUserClient(appId uint32, userId string) (client *Client) {

	manager.UserClientsLock.RLock()
	defer manager.UserClientsLock.RUnlock()

	userKey := GetUserKey(appId, userId)
	if value, ok := manager.UserClients[userKey]; ok {
		client = value
	}

	return
}

// 添加用户
func (manager *ClientManager) AddUsers(key string, client *Client) {
	manager.UserClientsLock.Lock()
	defer manager.UserClientsLock.Unlock()

	manager.UserClients[key] = client
}

// 删除用户
func (manager *ClientManager) DelUsers(key string) {
	manager.UserClientsLock.Lock()
	defer manager.UserClientsLock.Unlock()

	delete(manager.UserClients, key)
}

// 向全部成员(除了自己)发送数据
func (manager *ClientManager) sendAll(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore {
			conn.SendMsg(message)
		}
	}
}

// 用户建立连接事件
func (manager *ClientManager) EventSendUserMsg(message *msgs.SendUserMsg) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	fmt.Println(message.AppId)
	fmt.Println(message.UserId)
	client := Manager.GetUserClient(message.AppId, message.UserId)
	fmt.Println(client,"fffffff")
	fmt.Println(client.ToUid)
	fmt.Println(client.GroupsId)
	if client.ToUid=="" && client.GroupsId=="0"{

	}
	//data := msgs.GetTextMsgData(userId, msgId, message)
	//
	//// TODO::需要判断不在本机的情况
	//sendResults, err = SendUserMessageLocal(appId, userId, data)
	//if err != nil {
	//	fmt.Println("给用户发送消息", appId, userId, err)
	//}



	// client.Send <- []byte("连接成功")
}

// 用户建立连接事件
func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)
	fmt.Println("EventRegister 用户建立连接", client.Addr)
	// client.Send <- []byte("连接成功")
}

// 用户登录
func (manager *ClientManager) EventLogin(login *login) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()

	client := login.Client
	// 连接存在，在添加
	if _, ok := manager.Clients[login.Client]; ok {
		userKey := login.GetKey()
		manager.AddUsers(userKey, login.Client)
	}

	fmt.Println("EventLogin 用户登录", client.Addr, login.AppId, login.UserId)

	msgId := GetMsgIdTime()
	SendUserMessageAll(login.AppId, login.UserId, msgId, msgs.MessageActionEnter, "哈喽~")
}

// 用户断开连接
func (manager *ClientManager) EventUnregister(client *Client) {
	manager.DelClients(client)

	// 删除用户连接
	userKey := GetUserKey(client.AppId, client.UserId)
	manager.DelUsers(userKey)

	// 清除redis登录数据
	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err == nil {
		userOnline.LogOut()
		cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	}

	// 关闭 chan
	// close(client.Send)

	fmt.Println("EventUnregister 用户断开连接", client.Addr, client.AppId, client.UserId)

	if client.UserId != "" {
		msgId := GetMsgIdTime()
		_,_=SendUserMessageAll(client.AppId, client.UserId,msgId, msgs.MessageActionExit, "用户已经离开~")
	}
}

// 管道处理程序
func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			// 建立连接事件
			manager.EventRegister(conn)

		case login := <-manager.Login:
			// 用户登录
			manager.EventLogin(login)

		case conn := <-manager.Unregister:
			// 断开连接事件
			manager.EventUnregister(conn)

		case sendUserMsg := <-manager.SendUserMsg:
			//聊天室广播
			manager.EventSendUserMsg(sendUserMsg)
		}
	}
}

/**************************  manager info  ***************************************/
// 获取管理者信息
func GetManagerInfo(isDebug string) (managerInfo map[string]interface{}) {
	managerInfo = make(map[string]interface{})

	managerInfo["clientsLen"] = len(Manager.Clients)
	managerInfo["usersLen"] = len(Manager.UserClients)
	managerInfo["chanRegisterLen"] = len(Manager.Register)
	managerInfo["chanLoginLen"] = len(Manager.Login)
	managerInfo["chanUnregisterLen"] = len(Manager.Unregister)
	managerInfo["chanBroadcastLen"] = len(Manager.SendUserMsg)

	if isDebug == "true" {
		clients := make([]string, 0)
		for client := range Manager.Clients {
			clients = append(clients, client.Addr)
		}

		users := make([]string, 0)
		for key := range Manager.UserClients {
			users = append(users, key)
		}

		managerInfo["clients"] = clients
		managerInfo["users"] = users
	}

	return
}

// 获取用户所在的连接
func GetUserClient(appId uint32, userId string) (client *Client) {
	client = Manager.GetUserClient(appId, userId)

	return
}

// 定时清理超时连接
func ClearTimeoutConnections() {
	currentTime := uint64(time.Now().Unix())
	for client := range Manager.Clients {
		if client.IsHeartbeatTimeout(currentTime) {
			fmt.Println("心跳时间超时 关闭连接", client.Addr, client.UserId, client.LoginTime, client.HeartbeatTime)
			client.Socket.Close()
		}
	}
}

// 获取全部用户
func GetUserList() (userList []string) {

	userList = make([]string, 0)
	fmt.Println("获取全部用户")

	for _, v := range Manager.UserClients {
		userList = append(userList, v.UserId)
	}

	return
}

// 全员广播
func AllSendMessages(appId uint32, userId string, data string) {
	fmt.Println("全员广播", appId, userId, data)

	ignore := Manager.GetUserClient(appId, userId)
	Manager.sendAll([]byte(data), ignore)
}
