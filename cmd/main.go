package main

import (
	handler "controller/mod"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	homePage := handler.HomeHandler{}

	app.GET("/", homePage.HandleHomePage)
	app.GET("/loggedin", handler.HandleLogIn)
	app.GET("/loggedout", handler.HandleLogOut)
	app.Start(":8080")
}
