package router

import (
	"GoBlog/control"
	"github.com/labstack/echo"
)

func UserRouter(user *echo.Group) {
	user.GET("/page", control.UserPage)
	user.GET("/delete/:id", control.UserDelete)
	user.POST("", control.UserAdd)
	user.GET("/:id", control.UserGet)
	user.POST("/edit", control.UserEdit)
}
