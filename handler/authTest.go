package handler

import (
	"github.com/GarNepNep/MaxTrackr-go-app/model"
	"github.com/GarNepNep/MaxTrackr-go-app/view/page"
	"github.com/labstack/echo/v4"
)

func HandleAuthTest(c echo.Context) error {
	return page.ShowAuthTestPage(model.User{}).Render(c.Request().Context(), c.Response().Writer)
}
