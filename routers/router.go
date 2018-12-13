package routers

// @APIVersion 1.0.0
// @Title GASE Api
// @Description GASE autogenerate documents for your API

import (
	"GASE/controllers"
	"GASE/middlewares"

	"github.com/astaxie/beego"
)

func init() {

	middlewares.LoadMiddlewares()

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
			beego.NSRouter("/uuid/:cookie", &controllers.BriefsController{}, "get:GetOneByCookie"),
			beego.NSRouter("/:id", &controllers.BriefsController{}, "delete:Delete"),
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
		beego.NSNamespace("/projects",
			beego.NSInclude(
				&controllers.ProjectsController{},
			),
		),
		beego.NSNamespace("/attachments",
			beego.NSInclude(
				&controllers.AttachmentsController{},
			),
		),
		beego.NSNamespace("/sketchs",
			beego.NSInclude(
				&controllers.SketchsController{},
			),
		),
		beego.NSNamespace("/sketchs-files",
			beego.NSInclude(
				&controllers.SketchsFilesController{},
			),
		),
		beego.NSNamespace("/comments",
			beego.NSInclude(
				&controllers.CommentsController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
