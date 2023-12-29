package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/GarNepNep/MaxTrackr-go-app/handler"
)

func main() {
	app := echo.New()

	app.GET("/", handler.HandleHomePage)
	app.GET("/loggedin", handler.HandleLogIn)
	app.GET("/loggedout", handler.HandleLogOut)

	port := os.Getenv("PORT")
	fmt.Println("os.Getenv result", port)
	if port == "" {
		port = ":5000"
	}
	fmt.Println(port)
	app.Start(port)
}
