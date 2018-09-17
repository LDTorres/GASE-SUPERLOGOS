package main

import (
	_ "GASE/routers"

	_ "GASE/tasks"

	"github.com/astaxie/beego"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.BConfig.WebConfig.ViewsPath = "public/views"

	beego.SetStaticPath("admin/img", "public/assets/img")
	beego.SetStaticPath("admin/css", "public/assets/css")
	beego.SetStaticPath("admin/js", "public/assets/js")

	beego.Run()
}
