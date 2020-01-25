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
	err := DB.Select(&mods, "select * from user limit ?,?", pi*ps, ps)
	return mods, err
}

func UserCount() int {
	count := 0
	DB.Get(&count, "select count(id) as count from user")
	return count
}

func UserGet(id int64) (*User, error) {
	mod := &User{}
	err := DB.Get(mod, "select * from user where id = ? limit 1", id)
	return mod, err
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

func UserAdd(mod *User) error {
	tx, _ := DB.Begin()
	result, err := tx.Exec("insert into user (num, `name`, pass, phone, email, ctime, `status`) values (?,?,?,?,?,?,?)", mod.Num, mod.Name, mod.Pass, mod.Phone, mod.Email, mod.Ctime, mod.Status)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("row affected < 1")
	}
	tx.Commit()
	return nil
}

func UserExists(num string) bool {
	mod := User{}
	err := DB.Get(&mod, "select * from user where num = ?", num)
	if err != nil {
		return false
	}
	return true
}

func UserEdit(mod *User) error {
	tx, _ := DB.Beginx()
	result, err := tx.NamedExec("update user set `name`=:name, `phone`=:phone, `email`=:email where `id`=:id", mod)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("row affected < 1")
	}
	tx.Commit()
	return nil
}
