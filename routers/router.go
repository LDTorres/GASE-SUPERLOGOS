// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"GASE/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, controllers.GlobalMiddleware)

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/activities",
			beego.NSInclude(
				&controllers.ActivitiesController{},
			),
		),

		beego.NSNamespace("/carts",
			beego.NSInclude(
				&controllers.CartsController{},
			),
		),

		beego.NSNamespace("/clients",
			beego.NSInclude(
				&controllers.ClientsController{},
			),
		),

		beego.NSNamespace("/countries",
			beego.NSInclude(
				&controllers.CountriesController{},
			),
		),

		beego.NSNamespace("/coupons",
			beego.NSInclude(
				&controllers.CouponsController{},
			),
		),

		beego.NSNamespace("/currencies",
			beego.NSInclude(
				&controllers.CurrenciesController{},
			),
		),

		beego.NSNamespace("/gateways",
			beego.NSInclude(
				&controllers.GatewaysController{},
			),
		),

		beego.NSNamespace("/images",
			beego.NSInclude(
				&controllers.ImagesController{},
			),
		),

		beego.NSNamespace("/locations",
			beego.NSInclude(
				&controllers.LocationsController{},
			),
		),

		beego.NSNamespace("/orders",
			beego.NSInclude(
				&controllers.OrdersController{},
			),
		),

		beego.NSNamespace("/portfolios",
			beego.NSInclude(
				&controllers.PortfoliosController{},
			),
		),

		beego.NSNamespace("/prices",
			beego.NSInclude(
				&controllers.PricesController{},
			),
		),

		beego.NSNamespace("/sectors",
			beego.NSInclude(
				&controllers.SectorsController{},
			),
		),

		beego.NSNamespace("/services",
			beego.NSInclude(
				&controllers.ServicesController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSRouter("/", &controllers.UsersController{}, "post:Post"),
			beego.NSRouter("/", &controllers.UsersController{}, "get:GetAll"),
			beego.NSRouter("/:id", &controllers.UsersController{}, "get:Get"),
			beego.NSRouter("/:id", &controllers.UsersController{}, "put:Put"),
			beego.NSRouter("/:id", &controllers.UsersController{}, "delete:Delete"),
		),
	)
	beego.AddNamespace(ns)
}
