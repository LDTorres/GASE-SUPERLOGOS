package main

import (
	_ "./routers"

	"github.com/astaxie/beego"
)

func main() {

	beego.AppConfigPath = "conf/env.conf"

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
