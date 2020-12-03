package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//Db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       "root:123456@tcp(localhost:3306)/imdb?charset=utf8mb4&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	//	DefaultStringSize:         256,                                                                               // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
	//	DisableDatetimePrecision:  true,                                                                              // disable datetime precision support, which not supported before MySQL 5.6
	//	DontSupportRenameIndex:    true,                                                                              // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
	//	DontSupportRenameColumn:   true,                                                                              // use change when rename column, rename rename not supported before MySQL 8, MariaDB
	//	SkipInitializeWithVersion: false,                                                                             // smart configure based on used version
	//}), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize: 256, // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision: true, // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//设置连接池
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	fmt.Println("ok")
	fmt.Println(sqlDB)

}