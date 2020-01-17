package router

import (
	"GoBlog/control"
	"github.com/labstack/echo"
)

func ClassRouter(class *echo.Group) {
	class.GET("/all", control.ClassAll)
	class.GET("/page", control.ClassPage)
	class.POST("", control.ClassAdd)
	class.GET("/delete/:id", control.ClassDelete)
	class.GET("/:id", control.ClassGet)
	class.POST("/edit", control.ClassEdit)
}
