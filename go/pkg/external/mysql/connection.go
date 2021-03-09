package mysql

import (
	"fmt"
	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error
	var dsn string
	config ,err := ini.Load("./pkg/external/config/config.ini")
	passwordConf ,err := ini.Load("./pkg/external/config/password.ini")
	fmt.Print(config)
	user := config.Section("db").Key("MYSQL_USER").MustString("root")
	password := passwordConf.Section("db").Key("MYSQL_PASSWORD").MustString("hogehoge")
	hostName := config.Section("db").Key("MYSQL_HOSTNAME").MustString("localhost")
	port := config.Section("db").Key("MYSQL_PORT").String()
	dbName := config.Section("db").Key("MYSQL_DBNAME").String()
	dsn    = user + ":" + password + "@tcp(" + hostName + ":" + port + ")/" + dbName +"?charset=utf8&parseTime=True&loc=Local"
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