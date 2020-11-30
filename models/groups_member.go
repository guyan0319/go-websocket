package models

import (
	"time"
)

type GroupsMember struct {
	Id       int64     `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	GroupsId int64     `json:"groups_id" xorm:"not null default 0 comment('群id') BIGINT(20)"`
	Uid      int64     `json:"uid" xorm:"not null default 0 comment('用户id') index BIGINT(20)"`
	Type     int       `json:"type" xorm:"not null default 0 comment('类型（0:普通1群主）') TINYINT(4)"`
	Ctime    int       `json:"ctime" xorm:"not null default 0 comment('创建时间') INT(10)"`
	Mtime    time.Time `json:"mtime" xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
}
