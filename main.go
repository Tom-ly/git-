package main

import (
	"Work/db_mysql"
	_ "Work/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.ConnectDB()

	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()
	//gi


}

