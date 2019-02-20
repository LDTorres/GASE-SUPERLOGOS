package controllers

import (
	"GASE/models"
	"fmt"
)

// DatabaseController operations for Countries
type DatabaseController struct {
	BaseController
}

// URLMapping ...
func (c *DatabaseController) URLMapping() {
	c.Mapping("GenerateDatabase", c.GenerateDatabase)
}

// Get Generate Database ...
// @Title Generate Database
// @Description Generate Database
// @router /generate [get]
func (c *DatabaseController) GenerateDatabase() {
	// Add defaults to database
	count, _ := models.AddDefaultDataCurrencies()
	if count > 0 {
		fmt.Println("Added Currencies : ", count)
	}

	count, _ = models.AddDefaultDataCountries()
	if count > 0 {
		fmt.Println("Added Countries : ", count)
	}

	count, _ = models.AddDefaultDataSectors()
	if count > 0 {
		fmt.Println("Added Sectors : ", count)
	}

	/* count, _ = models.AddDefaultDataActivities()
	if count > 0 {
		fmt.Println("Added Activities : ", count)
	} */

	count, _ = models.AddDefaultDataGateways()
	if count > 0 {
		fmt.Println("Added Gateways : ", count)
	}

	count, _ = models.addRelationsGatewaysCurrencies()
	if count > 0 {
		fmt.Println("Added relations GatewaysCurrencies : ", count)
	}

	/*
		count, _ = models.AddDefaultDataServices()
		if count > 0 {
			fmt.Println("Added Services : ", count)
		}
	*/

	// models.AddDefaultDataPrices()

	models.AddDefaultDataUsers()

	c.Data["json"] = MessageResponse{
		Message:       "Generated Database",
		PrettyMessage: "Datos Generados",
	}

	c.ServeJSON()
}
