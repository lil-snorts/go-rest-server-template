package main

import (
	handler "controller/mod"

	"github.com/labstack/echo/v4"
)

var buttonAuth int

func main() {
	buttonAuth = 0
	app := echo.New()

	homePage := handler.HomeHandler{}

	app.GET("/", homePage.HandleHomePage)
	app.GET("/loggedin", handler.HandleLogIn)
	app.GET("/authTest", handler.HandleAuthTest)
	app.GET("/loggedout", handler.HandleLogOut)
	app.Start(":8080")
}
