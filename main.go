package main

import (
	"Go_Blog/models"
	_ "Go_Blog/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
}

func main() {
	//orm.debug = true
	beego.Run()
}
