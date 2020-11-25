package msgs

// 通用请求数据格式
type Request struct {
	Seq    string      `json:"seq"`            // 消息的唯一Id
	Action string      `json:"action"`         // 请求方法名
	Data   interface{} `json:"data,omitempty"` // 数据 json
}

// 接收用户登录请求数据
type Login struct {
	Token    string `json:"token,omitempty"` // token
	AppId    uint32 `json:"appId,omitempty"`
	UserId   string `json:"userId,omitempty"`
	ToUid    string `json:"toUid,omitempty"`
	GroupsId string `json:"groupsId,omitempty"`
}

// 心跳请求数据
type HeartBeat struct {
	UserId string `json:"userId,omitempty"`
}

//发消息数据
type SendUserMsg struct {
	AppId    uint32 `json:"appId,omitempty"`
	UserId   string `json:"userId,omitempty"`
	ToUid    string `json:"toUid,omitempty"`
	GroupsId string `json:"groupsId,omitempty"`
	MsgId string `json:"msgId,omitempty"`
	MsgType string `json:"msgType,omitempty"`//消息类型
    Message string `json:"message,omitempty"`
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
	ret = true

	return
}
