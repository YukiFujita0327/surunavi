package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error
	dsn := "root:@tcp(localhost:3306)/gorm_project?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// TODO create table gaetwaysを作ったら考える
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