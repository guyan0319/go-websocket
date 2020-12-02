package conf

import "gorm.io/driver/mysql"

var MysqlConfig = mysql.Config{
	DSN:                       "root:123456@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	DefaultStringSize:         256,                                                                               // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
	DisableDatetimePrecision:  true,                                                                              // disable datetime precision support, which not supported before MySQL 5.6
	DontSupportRenameIndex:    true,                                                                              // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
	DontSupportRenameColumn:   true,                                                                              // use change when rename column, rename rename not supported before MySQL 8, MariaDB
	SkipInitializeWithVersion: false,                                                                             // smart configure based on used version
}
