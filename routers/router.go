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
	"GASE/middlewares"

	"github.com/astaxie/beego"
)

func init() {

	middlewares.LoadFilters()

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

		beego.NSNamespace("/briefs",
			beego.NSRouter("/", &controllers.BriefsController{}, "post:Post"),
			beego.NSRouter("/", &controllers.BriefsController{}, "get:GetAll"),
			beego.NSRouter("/:id", &controllers.BriefsController{}, "get:Get"),
			beego.NSRouter("/:service/:cookie", &controllers.BriefsController{}, "get:GetOneByCookie"),
			beego.NSRouter("/:id", &controllers.BriefsController{}, "put:Put"),
			beego.NSRouter("/:id", &controllers.BriefsController{}, "delete:Delete"),
		),

		beego.NSNamespace("/forms",
			beego.NSRouter("/", &controllers.ServiceFormsController{}, "post:Post"),
			beego.NSRouter("/", &controllers.ServiceFormsController{}, "get:GetAll"),
			beego.NSRouter("/:id", &controllers.ServiceFormsController{}, "get:Get"),
			beego.NSRouter("/service/:slug", &controllers.ServiceFormsController{}, "get:GetOneByService"),
			beego.NSRouter("/:id", &controllers.ServiceFormsController{}, "put:Put"),
			beego.NSRouter("/:id", &controllers.ServiceFormsController{}, "delete:Delete"),
		),

		beego.NSNamespace("/payments-methods",
			beego.NSRouter("/", &controllers.PaymentsMethodsController{}, "post:Post"),
			beego.NSRouter("/", &controllers.PaymentsMethodsController{}, "get:GetAll"),
			beego.NSRouter("/:id", &controllers.PaymentsMethodsController{}, "get:Get"),
			beego.NSRouter("/:iso/:gateway", &controllers.PaymentsMethodsController{}, "get:GetOneByIsoAndGateway"),
			beego.NSRouter("/:id", &controllers.PaymentsMethodsController{}, "put:Put"),
			beego.NSRouter("/:id", &controllers.PaymentsMethodsController{}, "delete:Delete"),
		),

		beego.NSNamespace("/users",
			beego.NSRouter("/", &controllers.UsersController{}, "post:Post"),
			beego.NSRouter("/", &controllers.UsersController{}, "get:GetAll"),
			beego.NSRouter("/:id", &controllers.UsersController{}, "get:Get"),
			beego.NSRouter("/:id", &controllers.UsersController{}, "put:Put"),
			beego.NSRouter("/:id", &controllers.UsersController{}, "delete:Delete"),
			beego.NSRouter("/login", &controllers.UsersController{}, "post:Login"),
			beego.NSRouter("/change-password", &controllers.UsersController{}, "post:ChangePassword"),
		),
	)
	beego.AddNamespace(ns)
}
