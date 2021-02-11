package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"surunavi/go/pkg/adapter/gateways"
	"surunavi/go/pkg/domain"
	"surunavi/go/pkg/usecase"
	"surunavi/go/pkg/usecase/interfaces"
)

type LoginController struct {
	loginInteractor interfaces.LoginInteractor
}

func NewLoginController(conn *gorm.DB) *LoginController {
	return &LoginController{
		loginInteractor: &usecase.LoginInteractor{
			LoginRepository: &gateways.LoginRepository{
				Conn: conn,
			},
		},
	}
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
	controller.loginInteractor.Login(domain.UserInfo{Id: req.Id, Password: req.Password})
	return c.String(200, "hoge")
}
