package controllers

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"net/http"
	"surunavi/pkg/adapter/gateways"
	"surunavi/pkg/domain"
	"surunavi/pkg/usecase"
	_const "surunavi/pkg/usecase/const"
	"surunavi/pkg/usecase/interfaces"
)

type LoginController struct {
	loginInteractor interfaces.LoginInteractor
}

func NewLoginController(conn *gorm.DB) *LoginController {
	return &LoginController{
		loginInteractor: &usecase.LoginInteractor{
			LoginRepository: &gateways.LoginRepository{
				// DBのコネクションを貼りっぱなしになっているが、大丈夫か？
				Conn: conn,
			},
		},
	}
}

func (controller *LoginController) Login(c echo.Context) error {
	type(
		Request struct {
			UserId   string `json:"Id"`
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
	loginSuccess, userInfo, err := controller.loginInteractor.Login(domain.UserInfo{Id: req.UserId, Password: req.Password})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	res := Response{
		LoginSuccess: loginSuccess,
		UserId:       userInfo.Id,
		UserName:     userInfo.Name,
	}
	return c.JSON(http.StatusOK, res)
}
