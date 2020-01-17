package control

import (
	"GoBlog/common"
	"github.com/labstack/echo"
)

func Index(ctx echo.Context) error {
	println("???????????")
	return ctx.JSON(common.Succ("Success"))
}
