package handler

import (
	"github.com/GarNepNep/MaxTrackr-go-app/model"
	"github.com/GarNepNep/MaxTrackr-go-app/view/page"
	"github.com/labstack/echo/v4"
)

func HandleHomePage(c echo.Context) error {
	usr := model.User{
		Name: "yipee@yahoo.com"}
	return page.Show(usr).Render(c.Request().Context(), c.Response())
}
