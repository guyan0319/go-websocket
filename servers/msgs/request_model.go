package msgs

// 通用请求数据格式
type Request struct {
	Seq  string      `json:"seq"`            // 消息的唯一Id
	Cmd  string      `json:"cmd"`            // 请求命令字
	Data interface{} `json:"data,omitempty"` // 数据 json
}

// 接收用户登录请求数据
type Login struct {
	Token  string `json:"Token"` // token
	AppId  uint32 `json:"appId,omitempty"`
	UserId string `json:"userId,omitempty"`
}

// 心跳请求数据
type HeartBeat struct {
	UserId string `json:"userId,omitempty"`
}

//验证登录用户
func (l *Login) CheckLogin() (ret bool) {
	ret = true
	if l.UserId == "" || len(l.UserId) >= 20 {
		ret = false
		return
	}


	return
}

//验证token
func (l *Login) CheckToken() (ret bool) {
	ret=true

	return
}
