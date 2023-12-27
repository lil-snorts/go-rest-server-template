package handler

import (
	"fmt"

	"github.com/GarNepNep/MaxTrackr-go-app/model"
	"github.com/GarNepNep/MaxTrackr-go-app/view/page"
	"github.com/labstack/echo/v4"
)

func HandleLogIn(c echo.Context) error {
	fmt.Println(c, "\n", c.QueryParam("code"))
	return page.ShowLoggedInPage(model.User{}).Render(c.Request().Context(), c.Response().Writer)
}
