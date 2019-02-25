package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ActivitiesController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:AttachmentsController"],
		beego.ControllerComments{
			Method: "GetAttachmentsByUUID",
			Router: `/attachment/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:BriefsController"] = append(beego.GlobalControllerRouter["GASE/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:BriefsController"] = append(beego.GlobalControllerRouter["GASE/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:BriefsController"] = append(beego.GlobalControllerRouter["GASE/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:BriefsController"] = append(beego.GlobalControllerRouter["GASE/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:BriefsController"] = append(beego.GlobalControllerRouter["GASE/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "GetOneByCookie",
			Router: `/:service/:cookie`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:BriefsController"] = append(beego.GlobalControllerRouter["GASE/controllers:BriefsController"],
		beego.ControllerComments{
			Method: "GetImageByUUID",
			Router: `/image/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CartsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CartsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CartsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CartsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CartsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CartsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CartsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CartsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CartsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CartsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CartsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CartsController"],
		beego.ControllerComments{
			Method: "GetOneByCookie",
			Router: `/cookie`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CartsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CartsController"],
		beego.ControllerComments{
			Method: "AddOrDeleteServices",
			Router: `/services`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "ChangePasswordRequest",
			Router: `/change-password/:email`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "ChangePassword",
			Router: `/change-password/:token`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "GetOneByEmail",
			Router: `/email/:email`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ClientsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ClientsController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CommentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CommentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CommentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CommentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CommentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CommentsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "GetAttachmentsByUUID",
			Router: `/attachment/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CountriesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CountriesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CountriesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CountriesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CountriesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CountriesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CountriesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "GetOneByIso",
			Router: `/iso/:iso`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CountriesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CountriesController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CouponsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CouponsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CouponsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CouponsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CouponsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CouponsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CouponsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CouponsController"] = append(beego.GlobalControllerRouter["GASE/controllers:CouponsController"],
		beego.ControllerComments{
			Method: "GetOneByCode",
			Router: `/value/:code`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["GASE/controllers:CurrenciesController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:DatabaseController"] = append(beego.GlobalControllerRouter["GASE/controllers:DatabaseController"],
		beego.ControllerComments{
			Method: "GenerateDatabase",
			Router: `/generate`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "AddNewsCurrencies",
			Router: `/:id/currencies`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "DeleteCurrencies",
			Router: `/:id/currencies`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "GetSafetypayNotifications",
			Router: `/safetypay/notifications`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "GetSafetypayNotificationsConfirm",
			Router: `/safetypay/notifications/confirm`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "GetSafetypayTestRequestToken",
			Router: `/safetypay/test-express-token`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:GatewaysController"] = append(beego.GlobalControllerRouter["GASE/controllers:GatewaysController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ImagesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ImagesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ImagesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ImagesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ImagesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ImagesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ImagesController"],
		beego.ControllerComments{
			Method: "ServeImageBySlug",
			Router: `/slug/:slug`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LeadsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LeadsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LeadsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LeadsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LeadsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LeadsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LeadsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LeadsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LeadsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LeadsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LeadsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LeadsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LeadsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LeadsController"],
		beego.ControllerComments{
			Method: "Newsletter",
			Router: `/newsletter`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LeadsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LeadsController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LocationsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LocationsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LocationsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LocationsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LocationsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LocationsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:LocationsController"] = append(beego.GlobalControllerRouter["GASE/controllers:LocationsController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:OrdersController"] = append(beego.GlobalControllerRouter["GASE/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:OrdersController"] = append(beego.GlobalControllerRouter["GASE/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:OrdersController"] = append(beego.GlobalControllerRouter["GASE/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:OrdersController"] = append(beego.GlobalControllerRouter["GASE/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:OrdersController"] = append(beego.GlobalControllerRouter["GASE/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:OrdersController"] = append(beego.GlobalControllerRouter["GASE/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:OrdersController"] = append(beego.GlobalControllerRouter["GASE/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "GetSelf",
			Router: `/self`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:OrdersController"] = append(beego.GlobalControllerRouter["GASE/controllers:OrdersController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "GetByCustomSearch",
			Router: `/custom-search`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "GetPortfoliosByCountry",
			Router: `/iso`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "GetOneBySlug",
			Router: `/slug/:slug`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"] = append(beego.GlobalControllerRouter["GASE/controllers:PortfoliosController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PricesController"] = append(beego.GlobalControllerRouter["GASE/controllers:PricesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PricesController"] = append(beego.GlobalControllerRouter["GASE/controllers:PricesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PricesController"] = append(beego.GlobalControllerRouter["GASE/controllers:PricesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PricesController"] = append(beego.GlobalControllerRouter["GASE/controllers:PricesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:PricesController"] = append(beego.GlobalControllerRouter["GASE/controllers:PricesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ProjectsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ProjectsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ProjectsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ProjectsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ProjectsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ProjectsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ProjectsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ProjectsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ProjectsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ProjectsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ProjectsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ProjectsController"],
		beego.ControllerComments{
			Method: "GenerateUploadToken",
			Router: `/:id/token`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ProjectsController"] = append(beego.GlobalControllerRouter["GASE/controllers:ProjectsController"],
		beego.ControllerComments{
			Method: "VerifyUploadToken",
			Router: `/token/:token`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ReleasesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ReleasesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ReleasesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ReleasesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ReleasesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ReleasesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ReleasesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ReleasesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ReleasesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ReleasesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SectorsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SectorsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SectorsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SectorsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SectorsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SectorsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SectorsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SectorsController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ServicesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ServicesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ServicesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ServicesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ServicesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ServicesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "GetPricesServiceByCountry",
			Router: `/:id/prices`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ServicesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "RestoreFromTrash",
			Router: `/:id/restore`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:ServicesController"] = append(beego.GlobalControllerRouter["GASE/controllers:ServicesController"],
		beego.ControllerComments{
			Method: "GetAllFromTrash",
			Router: `/trashed`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsController"],
		beego.ControllerComments{
			Method: "NewPublicSketch",
			Router: `/token/:token`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"] = append(beego.GlobalControllerRouter["GASE/controllers:SketchsFilesController"],
		beego.ControllerComments{
			Method: "GetAttachmentsByUUID",
			Router: `/attachment/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:UsersController"] = append(beego.GlobalControllerRouter["GASE/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:UsersController"] = append(beego.GlobalControllerRouter["GASE/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:UsersController"] = append(beego.GlobalControllerRouter["GASE/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:UsersController"] = append(beego.GlobalControllerRouter["GASE/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:UsersController"] = append(beego.GlobalControllerRouter["GASE/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:UsersController"] = append(beego.GlobalControllerRouter["GASE/controllers:UsersController"],
		beego.ControllerComments{
			Method: "ChangePassword",
			Router: `/change-password`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["GASE/controllers:UsersController"] = append(beego.GlobalControllerRouter["GASE/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
