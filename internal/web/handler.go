package web

import (
	"github.com/labstack/echo/v4"
)

func (r *routing) Hello(ctx echo.Context) error {
	return ctx.JSON(200, map[string]string{"data": "hui"})
}
