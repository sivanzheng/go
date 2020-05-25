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

	article := app.Group("/article")
	ArticleRouter(article)

	app.Start(":9999")
}
