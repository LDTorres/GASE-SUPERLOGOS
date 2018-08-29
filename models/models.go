package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //
	"github.com/gosimple/slug"
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

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqlConnData.user+":"+mysqlConnData.pass+"@/"+mysqlConnData.dbName+"?charset=utf8")

	orm.RegisterModel(new(Activities), new(Clients), new(Countries), new(Coupons), new(Currencies), new(Gateways), new(Images), new(Locations), new(Orders), new(Portfolios), new(Prices), new(Sectors), new(Services))

	// Add defaults to database
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

//ValidateExists FK
func ValidateExists(modelName string, model interface{}) (err error) {

	switch modelName {
	case "Currency":
		o := orm.NewOrm()

		if err = o.Read(model.(*Currencies)); err != nil {
			return err
		}
	}

	return
}

//GenerateSlug return a slug
func GenerateSlug(modelName string, name string) (generatedSlug string) {

	o := orm.NewOrm()
	generatedSlug = slug.Make(name)
	var slugInt int

	switch modelName {
	case "Countries":

		var countries []*Countries
		o.QueryTable("countries").Filter("slug__startswith", generatedSlug).All(&countries)

		for _, val := range countries {
			formatSlug(val.Slug, generatedSlug, &slugInt)
		}
	}

	if slugInt > 0 {
		generatedSlug = generatedSlug + "-" + strconv.Itoa(slugInt+1)
	}

	return
}

func formatSlug(actualSlug string, generatedSlug string, slugInt *int) (err error) {

	replacedSlug := strings.Replace(actualSlug, generatedSlug, "", 1)

	replacedSlug = strings.Replace(replacedSlug, "-", "", 1)

	if replacedSlug == "" {
		replacedSlug = "1"
	}

	slugNumber, err := strconv.Atoi(replacedSlug)

	if err != nil {
		return
	}

	if slugNumber > *slugInt {
		*slugInt = slugNumber
	}

	return
}
