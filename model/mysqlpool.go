package model

import (
	"go-websocket/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)
var Db *gorm.DB
func init()  {
	var err error
	Db, err = gorm.Open(mysql.New(conf.MysqlConfig()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if conf.Cfg.ShowSql {
		Db = Db.Debug()
	}

	//设置连接池
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := Db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
