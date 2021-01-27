package main
import (
	"github.com/labstack/echo"
	"go-ca-webapi/surunavi/go/driver"
	"log"
	"os"
	"github.com/jinzhu/gorm"
)
func main() {
	dbConn, closeFunc, err := driver.NewDBConnection(
		os.Getenv("PRJ_GROUP_USERNAME"),
		os.Getenv("PRJ_GROUP_PASSWORD"),
		os.Getenv("PRJ_GROUP_INSTANCE"),
		os.Getenv("PRJ_GROUP_DBNAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	app := Initialize(dbConn, echo.New())
	app.Start()
}
