package models

import (
	"time"
)

type Message struct {
	Id       int64     `json:"id" xorm:"comment('ID') BIGINT(20)"`
	MsgSerId int64     `json:"msg_ser_id" xorm:"not null default 0 comment('服务端消息id') index BIGINT(20)"`
	GroupsId int       `json:"groups_id" xorm:"default 0 comment('群id') index INT(10)"`
	FromUid  int64     `json:"from_uid" xorm:"not null default 0 comment('发送方UID') index BIGINT(20)"`
	ToUid    int64     `json:"to_uid" xorm:"not null default 0 comment('接收方UID') index BIGINT(20)"`
	Content  string    `json:"content" xorm:"not null default '' comment('内容') VARCHAR(1000)"`
	Type     int       `json:"type" xorm:"not null default 0 comment('消息类型') TINYINT(1)"`
	Appname  int       `json:"appname" xorm:"not null default 0 comment('应用名称') TINYINT(3)"`
	Send     int       `json:"send" xorm:"not null default 0 comment('发送状态') TINYINT(1)"`
	Readed   int       `json:"readed" xorm:"not null default 0 comment('读取状态') TINYINT(1)"`
	Ext      string    `json:"ext" xorm:"default '' comment('扩展字段') VARCHAR(1024)"`
	State    int       `json:"state" xorm:"not null default 0 comment('状态') TINYINT(1)"`
	AdminId  int       `json:"admin_id" xorm:"not null default 0 comment('审核人') INT(10)"`
	Ctime    int       `json:"ctime" xorm:"not null default 0 comment('创建时间') INT(11)"`
	Mtime    time.Time `json:"mtime" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
}
