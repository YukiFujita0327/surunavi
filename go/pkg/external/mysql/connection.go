package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error
	var dsn string
	user := os.Getenv("GOROOT")
	password := os.Getenv("MYSQL_PASSWORD")
	hostname := os.Getenv("MYSQL_HOSTNAME")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DBNAME")
	dsn	= user + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + dbname +"?charset=utf8&parseTime=True&loc=Local"
 	fmt.Printf(user)
	fmt.Println(dsn)
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