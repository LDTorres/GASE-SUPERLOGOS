package main

import (
	_ "GASE/routers"

	"github.com/astaxie/beego"
)

func main() {

	//beego.AppConfig.Set("PATH", "conf/env.conf")

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
