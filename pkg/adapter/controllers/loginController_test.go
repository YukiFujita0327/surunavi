package controllers

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
	"testing"
)

func TestNewLoginController(t *testing.T) {
	con,_,_ := getDBMock()
	result := NewLoginController(con)
	fmt.Println(result.loginInteractor)
	fmt.Println(result.loginInteractor)
}

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	conn, err := gorm.Open("postgres", db)
	if err != nil {
		return nil, nil, err
	}
	return conn, mock, nil
}