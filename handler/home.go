package handler

import (
	pages "github.com/GarNepNep/MaxTrackr-go-app/view"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func (hh HomeHandler) HandleHomePage(c echo.Context) error {
	return pages.Show().Render(c.Request().Context(), c.Response())
}
