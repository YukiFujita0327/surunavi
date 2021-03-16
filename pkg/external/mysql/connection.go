package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error
	var dsn string
	DbConfig := NewDbConfig()
	dsn    = DbConfig.User + ":" + DbConfig.Password + "@tcp(" + DbConfig.HostName + ":" + DbConfig.Port + ")/" + DbConfig.DbName +"?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func CloseConn() {
	_db, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = _db.Close()
	if err != nil {
		panic(err)
	}
}