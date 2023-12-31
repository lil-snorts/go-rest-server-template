package main

import (
	echo "github.com/labstack/echo/v4"

	"maxtrackr/src/handler"
)

func main() {
	app := echo.New()

	app.GET("/", handler.HandleHomePage)
	app.GET("/loggedin", handler.HandleLogIn)
	app.GET("/loggedout", handler.HandleLogOut)
	app.Start(":5000")
}
