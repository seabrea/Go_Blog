package models

import (
	"time"
	"os"
	"path"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id         int64
	Title      string
	CreateTime time.Time `orm:"index"`
	Views      int64     `orm:"index"`
}

type Topic struct {
	Id         int64
	Uid        int64
	Title      string
	CreateTime time.Time `orm:"index"`
}

func RegisterDB(){
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)
}
