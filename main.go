package main

import (
	_ "GoBlog/routers"
	"GoBlog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.debug = true
	beego.Run()
}
