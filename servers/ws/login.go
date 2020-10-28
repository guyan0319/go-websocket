package ws

import "fmt"

// 用户登录
type login struct {
	AppId  uint32
	UserId string
	Client *Client
}

// 读取客户端数据
func (l *login) GetKey() (key string) {
	key = GetUserKey(l.AppId, l.UserId)
	return
}
// 获取用户key
func GetUserKey(appId uint32, userId string) (key string) {
	key = fmt.Sprintf("%d_%s", appId, userId)

	return
}
