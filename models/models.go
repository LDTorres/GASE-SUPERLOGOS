package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type mysqlConnData struct {
	user   string
	pass   string
	ip     string
	dbName string
}

func init() {

	RunMode := beego.BConfig.RunMode

	var mysqlConnData mysqlConnData

	mysqlConnData.user = beego.AppConfig.String(RunMode + "::mysqluser")
	mysqlConnData.pass = beego.AppConfig.String(RunMode + "::mysqlpass")
	mysqlConnData.dbName = beego.AppConfig.String(RunMode + "::mysqldb")

	//fmt.Println(mysqlConnData)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqlConnData.user+":"+mysqlConnData.pass+"@/"+mysqlConnData.dbName+"?charset=utf8")

	orm.RegisterModel(new(Activities), new(Clients), new(Countries), new(Coupons), new(Currencies), new(Gateways), new(Images), new(Locations), new(Orders), new(Portfolios), new(Prices), new(Sectors), new(Services))

	/* 	// Create database from models.
	   	name := "default"

	   	// Drop table and re-create.
	   	force := true

	   	// Print log.
	   	verbose := true

	   	// Error.
	   	err := orm.RunSyncdb(name, force, verbose)
	   	if err != nil {
	   		fmt.Println(err)
	   	} */

	count, _ := AddDefaultDataCurrencies()
	if count > 0 {
		fmt.Println("Added Currencies : ", count)
	}

	count, _ = addDefaultDataCountries()
	if count > 0 {
		fmt.Println("Added Countries : ", count)
	}

	count, _ = AddDefaultDataSectors()
	if count > 0 {
		fmt.Println("Added Sectors : ", count)
	}

	count, _ = addDefaultDataActivities()
	if count > 0 {
		fmt.Println("Added Activities : ", count)
	}

	count, _ = AddDefaultDataGateways()
	if count > 0 {
		fmt.Println("Added Gateways : ", count)
	}

	count, _ = addRelationsGatewaysCurrencies()
	if count > 0 {
		fmt.Println("Added relations GatewaysCurrencies : ", count)
	}

	count, _ = AddDefaultDataServices()
	if count > 0 {
		fmt.Println("Added Services : ", count)
	}

	count, _ = AddDefaultDataPrices()
	if count > 0 {
		fmt.Println("Added Prices : ", count)
	}
}
