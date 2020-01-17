package control

import (
	"GoBlog/common"
	"GoBlog/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"strconv"
	"time"
)

type login struct {
	Num  string
	Pass string
}

func UserLogin(ctx echo.Context) error {
	ipt := login{}
	err := ctx.Bind(&ipt)
	log.Println(ipt)
	if err != nil {
		return ctx.JSON(common.ErrIpt("Invalid Input", err.Error()))
	}
	mod, err := model.UserLogin(ipt.Num)
	if err != nil {
		return ctx.JSON(common.ErrIpt("Error User", err.Error()))
	}
	log.Println(mod.Pass)
	if mod.Pass != ipt.Pass {
		return ctx.JSON(common.ErrIpt("Error Password"))
	}

	claims := model.UserClaims{
		Id:   mod.Id,
		Num:  mod.Num,
		Name: mod.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(common.SECRET_KEY))
	if err != nil {
		fmt.Print(err)
	}
	return ctx.JSON(common.Succ("Success", ss))
}

type Page struct {
	Pi int `json:"pi"`
	Ps int `json:"ps"`
}

func UserPage(ctx echo.Context) error {
	ipt := Page{}
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(common.ErrIpt("Error Input", err.Error()))
	}
	if ipt.Ps < 1 || ipt.Ps > 50 {
		ipt.Ps = 10
	}
	if ipt.Pi < 1 {
		return ctx.JSON(common.ErrOpt("Error Input"))
	}
	count := model.UserCount()
	mods, err := model.UserPage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(common.ErrOpt("No Data", err.Error()))
	}
	return ctx.JSON(common.Page("", mods, count))
}

func UserDelete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(common.ErrIpt("Error Input", err.Error()))
	}
	uid, _ := ctx.Get("uid").(int64)
	if uid == id {
		return ctx.JSON(common.Fail("Can Not Delete Yourself"))
	}
	err = model.UserDelete(id)
	if err != nil {
		return ctx.JSON(common.Fail("Delete Failed", err.Error()))
	}
	return ctx.JSON(common.Succ("Delete Success"))

}
