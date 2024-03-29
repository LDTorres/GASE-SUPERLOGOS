package main

import (
	_ "GASE-SUPERLOGOS/routers"

	_ "GASE-SUPERLOGOS/tasks"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Country-Iso", "authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	/* beego.BConfig.ServerName = "liderlogo"
	beego.BConfig.WebConfig.FlashName = "LIDERLOGO_FLASH"
	beego.BConfig.Listen.ServerTimeOut = 60 */

	beego.Run()
}
