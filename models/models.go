package models

import (
	"os"
	"path"
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"

	// Register sqlite
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName        = "data/beeblog.db"
	sqlite3Driver = "sqlite3"
)

// Category 目录
type Category struct {
	ID         int64
	Title      string
	CreateTime time.Time `orm:"index"`
	Views      int64     `orm:"index"`
}

// Topic 文章
type Topic struct {
	ID         int64
	UID        int64
	Title      string
	CreateTime time.Time `orm:"index"`
}

func RegisterDB() {
	if !com.IsExist(dbName) {
		os.MkdirAll(path.Dir(dbName), os.ModePerm)
		os.Create(dbName)
	}

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDataBase("default", sqlite3Driver, dbName, 10)
}
