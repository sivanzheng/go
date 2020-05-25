package model

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type Class struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func ClassPage(pi, ps int) ([]Class, error) {
	mods := make([]Class, 0, ps)
	err := DB.Select(&mods, "select * from class limit ?,?", pi*ps, ps)
	return mods, err
}

func ClassCount() int {
	count := 0
	DB.Get(&count, "select count(id) as count from class")
	return count
}

func ClassAll() ([]Class, error) {
	mods := make([]Class, 0, 8)
	err := DB.Select(&mods, "select * from class")
	return mods, err
}

func ClassGet(id int64) (*Class, error) {
	mod := &Class{}
	err := DB.Get(mod, "select * from class where id =? limit 1", id)
	return mod, err
}

func ClassAdd(mod *Class) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("insert into class(`name`, `desc`) values (?,?)", mod.Name, mod.Desc)
	if err != nil {
		tx.Rollback()
		return nil
	}
	rows, err := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affected < 1")
	}
	tx.Commit()
	return nil

}

func ClassEdit(mod *Class) error {
	println(mod.Id, mod.Desc, mod.Name)
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("update class set name=?, `desc`=? where id =?", mod.Name, mod.Desc, mod.Id)
	if err != nil {
		tx.Rollback()
		return nil
	}
	rows, err := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affected < 1")
	}
	tx.Commit()
	return nil

}

func ClassDelete(id int64) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("delete from class where id =?", id)
	if err != nil {
		tx.Rollback()
		return nil
	}
	rows, err := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affected < 1")
	}
	tx.Commit()
	return nil
}

func ClassNameById(id int64) string {
	println(id)
	name := ""
	DB.Get(&name, "select name from class where id = ?", id)
	return name
}

func ClassNameByIds(ids []int64) map[int64]string {
	query, args, _ := sqlx.In("select * from class where id in (?)", ids)
	query = DB.Rebind(query)
	mods := make([]Class, 0, len(ids))
	DB.Select(&mods, query, args...)
	classNames := make(map[int64]string)
	for i := 0; i < len(mods); i++ {
		classNames[mods[i].Id] = mods[i].Name
	}
	println(&classNames)
	return classNames
}
