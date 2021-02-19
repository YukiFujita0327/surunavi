package controllers

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"net/http"
	"surunavi/go/pkg/adapter/gateways"
	"surunavi/go/pkg/domain"
	_const "surunavi/go/pkg/domain/const"
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
			LoginSuccess _const.LoginResultType `json:"LoginSuccess"`
			UserId       string                 `json:"UserId"`
			UserName     string                 `json:"UserName"`
		}
	)
	req := Request{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
		// TODO Interactor呼び出し + エラーハンドリングを書く
	loginSuccess,userinfo := controller.loginInteractor.Login(domain.UserInfo{Id: req.Id, Password: req.Password})
	res := Response{
		LoginSuccess: loginSuccess,
		UserId:       userinfo.Id,
		UserName:     userinfo.Name,
	}

	return c.JSON(http.StatusOK, res)
}
