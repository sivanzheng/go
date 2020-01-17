package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strings"
	"time"
)

type User struct {
	Id     int       `json:"id"`
	Num    string    `json:"num"`
	Name   string    `json:"name"`
	Pass   string    `json:"pass"`
	Phone  string    `json:"phone"`
	Email  string    `json:"email"`
	Status int       `json:"status"`
	Ctime  time.Time `json:"ctime"`
}

type UserClaims struct {
	Id   int    `json:"id"`
	Num  string `json:"num"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func UserLogin(num string) (User, error) {
	mod := User{}
	log.Println(num)
	err := DB.Get(&mod, "select * from user where num=? limit 1", strings.Trim(num, " "))
	return mod, err
}

func UserPage(pi, ps int) ([]User, error) {
	mods := make([]User, 0, ps)
	err := DB.Select(&mods, "select * from user limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

func UserCount() int {
	count := 0
	DB.Get(&count, "select count(id) as count from user")
	return count
}

func UserDelete(id int64) error {
	tx, _ := DB.Begin()
	result, err := tx.Exec("delete from user where id= ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affect < 1")
	}
	tx.Commit()
	return nil
}
