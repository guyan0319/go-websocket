package msgs

import "encoding/json"

/************************  响应数据  **************************/
type Head struct {
	Seq      string    `json:"seq"`      // 消息的Id
	Action   string    `json:"action"`   // 请求方法名
	Response *Response `json:"response"` // 消息体
}

type Response struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"` // 数据 json
}

// push 数据结构体
type PushMsg struct {
	Seq  string `json:"seq"`
	Uid uint64 `json:"uid"`
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

// 设置返回消息
func NewResponseHead(seq string, action string, code uint32, codeMsg string, data interface{}) *Head {
	response := NewResponse(code, codeMsg, data)

	return &Head{Seq: seq, Action: action, Response: response}
}

func (h *Head) String() (headStr string) {
	headBytes, _ := json.Marshal(h)
	headStr = string(headBytes)

	return
}

func NewResponse(code uint32, msg string, data interface{}) *Response {
	return &Response{Code: code, Msg: msg, Data: data}
}
func (r *Response) String() (headStr string) {
	headBytes, _ := json.Marshal(r)
	headStr = string(headBytes)

	return
}
