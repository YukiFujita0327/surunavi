package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type LoginController struct {
	hoge string
}

func NewLoginController(conn *gorm.DB) *LoginController {
	return &LoginController{"hoge"}
}

func (controller *LoginController) Login(c echo.Context) error {
	type(
		Request struct {
			Id       string `json:"Id"`
			Password string `json:"Password"`
		}
		Response struct {
			UserInfo bool `json:"UserInfo"`
		}
	)
	req := Request{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	// TODO Interactor呼び出し + エラーハンドリングを書く
	return c.String(200, "hoge")
}
