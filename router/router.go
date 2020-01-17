package router

import (
	"GoBlog/control"
	"github.com/labstack/echo"
)

func Run() {
	app := echo.New()
	//app.Static("/static", "static")
	app.POST("/login", control.UserLogin)

	class := app.Group("/class")
	ClassRouter(class)

	user := app.Group("/user", ServerHeader) // 限制token
	UserRouter(user)

	app.Start(":9999")
}
