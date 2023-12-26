package handler

import (
	"github.com/GarNepNep/MaxTrackr-go-app/model"
	"github.com/GarNepNep/MaxTrackr-go-app/view/page"
	"github.com/labstack/echo/v4"
)

func HandleLogOut(c echo.Context) error {
	return page.ShowLoggedInPage(model.User{}).Render(c.Request().Context(), c.Response().Writer)
}
