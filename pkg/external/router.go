package external

import (
	"github.com/labstack/echo"
	"surunavi/pkg/adapter/controllers"
	"surunavi/pkg/external/mysql"
)

func Run() {
	defer mysql.CloseConn()
	conn := mysql.Connect()
	var e = echo.New()
	loginController := controllers.NewLoginController(conn)

	e.POST("/login", func(c echo.Context) error { return loginController.Login(c)})
	e.Start(":8080")
}
