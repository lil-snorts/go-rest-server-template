package main

import (
	handler "controller/mod"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	homePage := handler.HomeHandler{}

	app.GET("/", homePage.HandleHomePage)
	app.Start(":8080")
}
