package models

import (
	"time"
)

type Member struct {
	Id    int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Name  string    `json:"name" xorm:"not null default '' comment('姓名') VARCHAR(50)"`
	State int       `json:"state" xorm:"not null default 0 comment('状态(1:有效，0无效，9删除)') TINYINT(4)"`
	Ctime int       `json:"ctime" xorm:"not null default 0 comment('创建时间') INT(10)"`
	Mtime time.Time `json:"mtime" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
