package control

import (
	"GoBlog/common"
	"GoBlog/model"
	"github.com/labstack/echo"
	"strconv"
)

func ClassGet(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(common.ErrIpt("Invaild Input", err.Error()))
	}
	mod, err := model.ClassGet(id)
	if err != nil {
		return ctx.JSON(common.ErrOpt("No Data", err.Error()))
	}
	return ctx.JSON(common.Succ("Success", mod))
}

func ClassAll(ctx echo.Context) error {
	mods, err := model.ClassAll()
	if err != nil {
		ctx.JSON(common.Fail("No Data", err.Error()))
	}
	return ctx.JSON(common.Succ("", mods))
}

func ClassPage(ctx echo.Context) error {
	pi, err := strconv.Atoi(ctx.FormValue("pi"))
	if err != nil || pi < 0 {
		return ctx.JSON(common.Fail("Bad Request", err.Error()))
	}
	ps, err := strconv.Atoi(ctx.FormValue("ps"))
	if err != nil || ps > 50 || ps < 1 {
		return ctx.JSON(common.Fail("Bad Request", err.Error()))
	}

	count := model.ClassCount()
	if count < 1 {
		return ctx.JSON(common.ErrOpt("No Data"))
	}
	mods, err := model.ClassPage(pi, ps)
	if err != nil {
		return ctx.JSON(common.ErrOpt("No Data", err.Error()))
	}
	return ctx.JSON(common.Page("", mods, count))
}

func ClassAdd(ctx echo.Context) error {
	ipt := &model.Class{}
	ctx.Bind(ipt)

	if ipt.Name == "" {
		return ctx.JSON(common.ErrIpt("No Empty Name"))
	}
	if ipt.Desc == "" {
		return ctx.JSON(common.ErrIpt("No Empty Desc"))
	}

	err := model.ClassAdd(ipt)
	if err != nil {
		return ctx.JSON(common.Fail("Add Failed"))
	}
	return ctx.JSON(common.Succ("Add Success"))
}

func ClassEdit(ctx echo.Context) error {
	ipt := &model.Class{}
	ctx.Bind(ipt)

	if ipt.Id == 0 {
		return ctx.JSON(common.ErrIpt("No ID"))
	}
	if ipt.Name == "" {
		return ctx.JSON(common.ErrIpt("No Empty Name"))
	}
	if ipt.Desc == "" {
		return ctx.JSON(common.ErrIpt("No Empty Desc"))
	}

	err := model.ClassEdit(ipt)
	if err != nil {
		return ctx.JSON(common.Fail("Edit Failed"))
	}
	return ctx.JSON(common.Succ("Edit Success"))
}

func ClassDelete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(common.ErrIpt("Invaild Input", err.Error()))
	}
	err = model.ClassDelete(id)
	if err != nil {
		return ctx.JSON(common.Fail("Delete Failed", err.Error()))
	}
	return ctx.JSON(common.Succ("Delete Failed"))
}
