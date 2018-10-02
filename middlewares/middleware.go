package middlewares

import (
	"GASE/controllers"
	"errors"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//InsertMiddleware ...
func InsertMiddleware(ctrlName string, url string, pattern string, users []string) {

	beego.InsertFilter(url, beego.BeforeRouter, Middleware(ctrlName, pattern, users))

}

//LoadMiddlewares ...
func LoadMiddlewares() {
	for _, ctrlName := range ControllersNames {
		for pattern, users := range GetURLMapping(ctrlName) {
			if strings.Contains(pattern, ";") {
				methods := strings.Split(pattern, ";")

				InsertMiddleware(ctrlName, "/*/"+ctrlName+methods[0], pattern, users)
			} else {

				InsertMiddleware(ctrlName, "/*/"+ctrlName+pattern, pattern, users)
			}
		}
	}
}

//Middleware ...
func Middleware(controller string, pattern string, userTypes []string) func(ctx *context.Context) {

	return func(ctx *context.Context) {

		token := ctx.Input.Header("Authorization")

		excludeUrls := []string{
			"login", "carts", "change-password", "custom-search",
		}

		verifyToken := true

		// If the url is a excluded url then dont verify the token
		for _, excludeURL := range excludeUrls {
			if strings.Contains(ctx.Input.URL(), excludeURL) {
				verifyToken = false

				//If is a optional method like carts
				switch ctx.Input.Method() {
				case "DELETE":
					if excludeURL == "carts" {
						verifyToken = true
						break
					}
				case "GET":
					if excludeURL == "carts" {
						verifyToken = true
						break
					}
				}
				break
			}
		}

		// Return middleware
		if !verifyToken {
			beego.Debug("Exclude Url: Dont verify token")
			return
		}

		/* DEBUG */
		/* beego.Debug("Url: ", controller, "| Pattern: ", pattern)
		beego.Debug("Method: ", ctx.Input.Method())
		beego.Debug("Allowed users: ", userTypes) */

		// Verify global methods
		methods := []string{
			"POST", "DELETE", "PUT", "OPTIONS",
		}

		for _, method := range methods {
			if ctx.Input.Method() == method {
				if method == "OPTIONS" {
					beego.Debug("OPTIONS METHOD: Dont verify token")
					return
				}

				// Deny Access if the token is empty
				if token == "" {
					err := errors.New("Token No enviado")
					DenyAccess(ctx, err)
					return
				}
				_, err := controllers.VerifyToken(token, "Admin")

				if err != nil {
					beego.Debug("Invalid User Type in methods")
					err := errors.New("Usuario Invalido")
					DenyAccess(ctx, err)
					return
				}
			}
		}

		// Verify custom validation
		if strings.Contains(pattern, ";") {
			urlMapping := GetURLMapping(controller)
			splitted := strings.Split(pattern, ";")
			methods := splitted[1]

			if strings.Contains(methods, ctx.Input.Method()) {
				userTypes = urlMapping[pattern]
				beego.Debug("Allowed users after pattern: ", userTypes)
			}
		}

		// If Usertype is Guest
		if userTypes[0] == "Guest" {
			return
		}

		// Deny Access By Default
		denyAccess := true

		// Deny Access if the token is empty
		if token == "" {
			err := errors.New("Token No enviado")
			DenyAccess(ctx, err)
			return
		}

		for _, userType := range userTypes {
			user, _ := controllers.VerifyToken(token, userType)

			if userType == user.Type {
				denyAccess = false
				break
			}
		}

		if denyAccess {
			beego.Debug("Invalid User Type in types")
			err := errors.New("Usuario Invalido")
			DenyAccess(ctx, err)
		}
	}
}
