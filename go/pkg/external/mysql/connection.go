package mysql

import "github.com/jinzhu/gorm"

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", "root:@tcp(db:3306)/hoge")
	if err != nil {
		panic(err)
	}
	// TODO create table gaetwaysを作ったら考える
	return db
}

func CloseConn() {
	db.Close()
}