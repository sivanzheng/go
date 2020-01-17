package router

import (
	"GoBlog/common"
	"GoBlog/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderServer, "Echo/999")
		tokenString := ctx.Request().Header.Get("Authorization")
		println(tokenString)
		claims := model.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(common.SECRET_KEY), nil
		})
		println(token.Valid, err)
		if err == nil && token.Valid {
			println("next")
			ctx.Set("uid", claims.Id)
			return next(ctx)
		} else {
			return ctx.JSON(common.ErrJwt("invalid token"))
		}
	}
}
