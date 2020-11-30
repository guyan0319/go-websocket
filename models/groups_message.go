package models

import (
	"time"
)

type GroupsMessage struct {
	Id       int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	MsgSerId int64     `json:"msg_ser_id" xorm:"not null default 0 comment('服务端消息id') index(msg_ser_id) BIGINT(20)"`
	Uid      int64     `json:"uid" xorm:"not null default 0 comment('用户id') index(msg_ser_id) BIGINT(20)"`
	Readed   int       `json:"readed" xorm:"not null default 0 comment('读取状态（0:未读,1:已读）') TINYINT(4)"`
	Mtime    time.Time `json:"mtime" xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	Ctime    int       `json:"ctime" xorm:"not null default 0 comment('创建时间') INT(10)"`
}
