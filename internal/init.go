package internal

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

//Init initialize  the web server
func Init(){
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}