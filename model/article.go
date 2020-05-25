package model

import (
	"errors"
	"time"
)

type Article struct {
	Id        int64     `json:"id"`
	Cid       int64     `json:"cid"`
	ClassName string    `json:"class_name" db:"-"`
	Uid       int64     `json:"uid"`
	Title     string    `json:"title"`
	Origin    string    `json:"origin"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Hits      int64     `json:"hits"`
	Utime     time.Time `json:"utime"`
	Ctime     time.Time `json:"ctime"`
}

func ArticleCount() int {
	count := 0
	DB.Get(&count, "select count(*) from article")
	return count
}

func ArticlePage(pi, ps int) ([]Article, error) {
	mods := make([]Article, 0, ps)
	err := DB.Select(&mods, "select  * from article limit ?,?", pi*ps, ps)
	cids := make([]int64, 0, ps)
	for i := 0; i < len(mods); i++ {
		if !inOf(mods[i].Cid, cids) {
			cids = append(cids, mods[i].Cid)
		}
	}
	classNames := ClassNameByIds(cids)
	println(classNames)
	for i := 0; i < len(mods); i++ {
		mods[i].ClassName = classNames[mods[i].Cid]
		mods[i].Content = ""
	}
	return mods, err
}

func inOf(dst int64, arr []int64) bool {
	for i := 0; i < len(arr); i++ {
		if dst == arr[i] {
			return true
		}
	}
	return false
}

func ArticleDelete(id int64) error {
	tx, _ := DB.Beginx()
	result, err := tx.Exec("delete from article where id =?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affected < 1")
	}
	tx.Commit()
	return nil
}

func ArticleGet(id int64) (*Article, error) {
	mod := &Article{}
	err := DB.Get(mod, "select * from article where id =? limit 1", id)
	return mod, err
}

func ArticleAdd(mod *Article) error {
	tx, _ := DB.Beginx()
	result, err := tx.NamedExec("insert into article (`cid`, `uid`, `title`, `origin`, `author`, `content`, `hits`, `utime`, `ctime`) values (:cid, :uid, :title, :origin, :author, :content, :hits, :utime, :ctime)", mod)
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

func ArticleEdit(mod *Article) error {
	tx, _ := DB.Beginx()
	result, err := tx.NamedExec("update article set `cid`=:cid, `uid`=:uid, `title`=:title, `origin`=:origin, `author`=:author, `content`=:content, `hits`=:hits, `utime`=:utime, `ctime`=:ctime where `id`=:id ", mod)
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
