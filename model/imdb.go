package model

import (
	"time"
)

// Groups [...]
type Groups struct {
	ID    int64     `gorm:"primaryKey;column:id;type:bigint(20);not null"`   // 主键
	Name  string    `gorm:"column:name;type:varchar(50);not null"`           // 名称
	Appid int8      `gorm:"column:appid;type:tinyint(10) unsigned;not null"` // 应用id
	State int8      `gorm:"column:state;type:tinyint(4);not null"`           // 状态(1:有效，0无效，9删除)
	Ctime int       `gorm:"column:ctime;type:int(10);not null"`              // 创建时间
	Mtime time.Time `gorm:"column:mtime;type:timestamp;not null"`            // 修改时间
}

// GroupsMember [...]
type GroupsMember struct {
	ID       int64     `gorm:"primaryKey;column:id;type:bigint(20);not null"` // 主键
	GroupsID int64     `gorm:"column:groups_id;type:bigint(20);not null"`     // 群id
	UId      int64     `gorm:"index:uid;column:uid;type:bigint(20);not null"` // 用户id
	Type     int8      `gorm:"column:type;type:tinyint(4);not null"`          // 类型（0:普通1群主）
	Ctime    int       `gorm:"column:ctime;type:int(10);not null"`            // 创建时间
	Mtime    time.Time `gorm:"column:mtime;type:timestamp;not null"`          // 修改时间
}

// GroupsMessage [...]
type GroupsMessage struct {
	ID       int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	MsgSerID int64     `gorm:"index:msg_ser_id;column:msg_ser_id;type:bigint(20);not null"` // 服务端消息id
	UId      int64     `gorm:"index:msg_ser_id;column:uid;type:bigint(20);not null"`        // 用户id
	Readed   int8      `gorm:"column:readed;type:tinyint(4);not null"`                      // 读取状态（0:未读,1:已读）
	Mtime    time.Time `gorm:"column:mtime;type:timestamp;not null"`                        // 修改时间
	Ctime    int       `gorm:"column:ctime;type:int(10);not null"`                          // 创建时间
}

// Member [...]
type Member struct {
	ID    int64     `gorm:"primaryKey;column:id;type:bigint(20);not null"`
	Name  string    `gorm:"column:name;type:varchar(50);not null"` // 姓名
	State int8      `gorm:"column:state;type:tinyint(4);not null"` // 状态(1:有效，0无效，9删除)
	Ctime int       `gorm:"column:ctime;type:int(10);not null"`    // 创建时间
	Mtime time.Time `gorm:"column:mtime;type:timestamp;not null"`
}

// Message 消息表
type Message struct {
	ID       uint64    `gorm:"column:id;type:bigint(20) unsigned;not null"`                          // ID
	MsgSerID int64     `gorm:"index:msg_ser_id;column:msg_ser_id;type:bigint(20) unsigned;not null"` // 服务端消息id
	GroupsID int       `gorm:"index:groups_id;column:groups_id;type:int(10)"`                        // 群id
	FromUId  uint64    `gorm:"index:from_uid;column:from_uid;type:bigint(20) unsigned;not null"`     // 发送方UID
	ToUId    uint64    `gorm:"index:to_uid;column:to_uid;type:bigint(20) unsigned;not null"`         // 接收方UID
	Content  string    `gorm:"column:content;type:varchar(1000);not null"`                           // 内容
	Type     bool      `gorm:"column:type;type:tinyint(1) unsigned;not null"`                        // 消息类型
	Appname  uint8     `gorm:"column:appname;type:tinyint(3) unsigned;not null"`                     // 应用名称
	Send     bool      `gorm:"column:send;type:tinyint(1) unsigned;not null"`                        // 发送状态
	Readed   bool      `gorm:"column:readed;type:tinyint(1) unsigned;not null"`                      // 读取状态
	Ext      string    `gorm:"column:ext;type:varchar(1024)"`                                        // 扩展字段
	State    bool      `gorm:"column:state;type:tinyint(1) unsigned;not null"`                       // 状态
	AdminID  uint32    `gorm:"column:admin_id;type:int(10) unsigned;not null"`                       // 审核人
	Ctime    uint32    `gorm:"column:ctime;type:int(11) unsigned;not null"`                          // 创建时间
	Mtime    time.Time `gorm:"column:mtime;type:timestamp;not null"`                                 // 更新时间
}
