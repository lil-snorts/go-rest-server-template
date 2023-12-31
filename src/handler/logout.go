package handler

import (
	echo "github.com/labstack/echo/v4"

	"maxtrackr/src/model"
	"maxtrackr/src/view/page"
)

func HandleLogOut(c echo.Context) error {
	return page.ShowLoggedInPage(model.User{}).Render(c.Request().Context(), c.Response().Writer)
}
