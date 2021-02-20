package external

import (
	"github.com/labstack/echo"
	"surunavi/go/pkg/adapter/controllers"
	"surunavi/go/pkg/external/mysql"
)

func Run() {
	defer mysql.CloseConn()
	conn := mysql.Connect()
	var e = echo.New()
	loginController := controllers.NewLoginController(conn)

	e.GET("/login", func(c echo.Context) error { return loginController.Login(c)})
	e.Start(":8080")
}
