package router

import (
	"GoBlog/control"
	"github.com/labstack/echo"
)

func ArticleRouter(article *echo.Group) {
	article.GET("/page", control.ArticlePage)
	article.GET("/:id", control.ArticleGet)
	article.GET("/delete/:id", control.ArticleDelete)
	article.POST("/add", control.ArticleAdd)
	article.POST("/edit", control.ArticleEdit)
}
