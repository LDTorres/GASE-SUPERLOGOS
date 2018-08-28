package controllers

import (
	"github.com/astaxie/beego"

	"github.com/go-sql-driver/mysql"
)

// BaseController operations for Activities
type BaseController struct {
	beego.Controller
}

//ErrorHandler =
type ErrorHandler struct {
	Message       string `json:"message"`
	Code          uint16 `json:"code"`
	PrettyMessage string `json:"pretty_message"`
}

//ServeErrorJSON : Serve Json error
func (c *BaseController) ServeErrorJSON(err error) {

	if driverErr, ok := err.(*mysql.MySQLError); ok {

		switch driverErr.Number {
		case 1062:
			c.Ctx.Output.SetStatus(409)
			c.Data["json"] = ErrorHandler{
				Message:       "The element already exists",
				Code:          driverErr.Number,
				PrettyMessage: "El elemento de la base de datos ya existe",
			}
		default:
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = ErrorHandler{
				Message:       "An error has ocurred",
				Code:          driverErr.Number,
				PrettyMessage: "Un error ha ocurrido",
			}
		}
	}

	c.ServeJSON()
}

//BadRequest =
func (c *BaseController) BadRequest() {
	c.Ctx.Output.SetStatus(403)
	c.Data["json"] = ErrorHandler{
		Message:       "Bad request body",
		PrettyMessage: "Peticion mal formada",
	}
	c.ServeJSON()
}
