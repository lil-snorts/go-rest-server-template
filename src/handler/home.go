package handler

import (
	echo "github.com/labstack/echo/v4"

	"maxtrackr/src/model"
	"maxtrackr/src/view/page"
)

func HandleHomePage(c echo.Context) error {
	usr := model.User{
		Name: "yipee@yahoo.com"}
	return page.Show(usr).Render(c.Request().Context(), c.Response())
}
