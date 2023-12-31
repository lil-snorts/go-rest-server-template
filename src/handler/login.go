package handler

import (
	"fmt"

	echo "github.com/labstack/echo/v4"

	"maxtrackr/src/model"
	"maxtrackr/src/view/page"
)

func HandleLogIn(c echo.Context) error {
	fmt.Println(c, "\n", c.QueryParam("code"))
	return page.ShowLoggedInPage(model.User{}).Render(c.Request().Context(), c.Response().Writer)
}
