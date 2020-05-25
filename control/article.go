package control

import (
	"GoBlog/common"
	"GoBlog/model"
	"github.com/labstack/echo"
	"strconv"
	"time"
)

func ArticlePage(ctx echo.Context) error {
	ipt := Page{}
	err := ctx.Bind(&ipt)
	if err != nil {
		ctx.JSON(common.ErrIpt("Invaild Data"))
	}
	if ipt.Pi < 0 {
		ipt.Pi = 0
	}
	if ipt.Ps < 1 || ipt.Ps > 50 {
		ipt.Ps = 6
	}
	count := model.ArticleCount()
	if count < 1 {
		return ctx.JSON(common.ErrOpt("No Data"))
	}
	mods, err := model.ArticlePage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(common.ErrOpt("No Data", err.Error()))
	}
	return ctx.JSON(common.Page("Articles", mods, count))

}

func ArticleGet(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(common.ErrIpt("Invaild Data", err.Error()))
	}
	mod, err := model.ArticleGet(id)
	if err != nil {
		return ctx.JSON(common.ErrOpt("No Data", err.Error()))
	}
	return ctx.JSON(common.Succ("Success", mod))

}

func ArticleDelete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(common.ErrIpt("Invaild Data", err.Error()))
	}
	err = model.ArticleDelete(id)
	if err != nil {
		return ctx.JSON(common.Fail("Delete Fail", err.Error()))
	}
	return ctx.JSON(common.Succ("Delete Success"))
}

func ArticleAdd(ctx echo.Context) error {
	ipt := model.Article{}
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(common.ErrIpt("Invaild Data", err.Error()))
	}

	ipt.Utime = time.Now()
	ipt.Ctime = time.Now()
	err = model.ArticleAdd(&ipt)
	if err != nil {
		return ctx.JSON(common.Fail("Add Fail", err.Error()))
	}
	return ctx.JSON(common.Succ("Add Success"))
}

func ArticleEdit(ctx echo.Context) error {
	ipt := model.Article{}
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(common.ErrIpt("Invaild Data", err.Error()))
	}
	ipt.Utime = time.Now()
	err = model.ArticleEdit(&ipt)
	if err != nil {
		return ctx.JSON(common.Fail("Edit Failed", err.Error()))
	}
	return ctx.JSON(common.Succ("Edit Success"))
}
