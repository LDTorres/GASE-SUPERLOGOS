package controllers

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
)

// MiddlewareController operations for Middleware
type MiddlewareController struct {
	BaseController
}

// LoginToken =
type LoginToken struct {
	Type string `json:"tipo"`
	ID   string `json:"id"`
	jwt.StandardClaims
}

//VerifyToken =
func VerifyToken(tokenString string, userType string) (decodedToken *LoginToken, err error) {

	if tokenString == "" {
		return nil, errors.New("Token Vacio")
	}

	hmacSampleSecret := []byte("bazam")

	token, err := jwt.ParseWithClaims(tokenString, &LoginToken{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*LoginToken)

	if !ok && !token.Valid {
		return nil, err
	}

	//fmt.Println(tokenString, userType)
	//fmt.Println(claims)

	if claims.Type != userType {
		return nil, errors.New("Invalid User")
	}

	return claims, nil
}

// GenerateToken =
func (c *BaseController) GenerateToken(userType string, id string) (token string) {

	hmacSampleSecret := []byte("bazam")

	now := time.Now()

	// Create the Claims
	claims := LoginToken{
		userType,
		id,
		jwt.StandardClaims{
			ExpiresAt: now.AddDate(0, 0, 7).Unix(),
			Issuer:    "test",
		},
	}

	var newToken *jwt.Token
	newToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := newToken.SignedString(hmacSampleSecret)

	if err != nil {
		c.BadRequest(err)
	}

	return token
}

var (
	//ControllersNames ...
	ControllersNames = []string{
		"portfolios", "activities", "carts", "clients", "countries", "coupons", "currencies", "gateways", "images", "locations", "orders", "prices", "sectors", "services", "email", "briefs", "users",
	}
)

//GetURLMapping =
func GetURLMapping(route string) (validation map[string][]string) {

	portfolios := map[string][]string{
		";POST,PUT":       {"Admin"},
		"/:id;PUT,DELETE": {"Admin"},
	}

	activities := map[string][]string{
		"/:id;GET": {"Admin"},
	}

	validations := map[string]map[string][]string{
		"portfolios": portfolios,
		"activities": activities,
	}

	return validations[route]
}

//LoadFilters ...
func LoadFilters() {
	for _, ctrlName := range ControllersNames {
		for url, users := range GetURLMapping(ctrlName) {
			if strings.Contains(url, ";") {
				urlWithMethod := strings.Split(url, ";")
				beego.InsertFilter("/*/"+ctrlName+urlWithMethod[0], beego.BeforeRouter, Middleware(ctrlName, url, users))
			} else {
				beego.InsertFilter("/*/"+ctrlName+url, beego.BeforeRouter, Middleware(ctrlName, url, users))
			}
		}
	}
}

//Middleware =
func Middleware(controller string, pattern string, userTypes []string) func(ctx *context.Context) {

	return func(ctx *context.Context) {

		if userTypes[0] == "Guest" {
			return
		}

		urlMapping := GetURLMapping(controller)

		token := ctx.Input.Header("Authorization")

		if token == "" {
			err := errors.New("Token Invalido")
			DenyPermision(ctx, err)
			return
		}

		beego.Debug("Url: ", controller, "| Match: ", pattern)
		beego.Debug("Allowed users: ", userTypes)
		beego.Debug("Method: ", ctx.Input.Method())

		valid := false

		// if has methods
		if strings.Contains(pattern, ";") {
			splitted := strings.Split(pattern, ";")
			methods := splitted[1]

			if strings.Contains(methods, ctx.Input.Method()) {

				userTypes = urlMapping[pattern]

				beego.Debug("Allowed users in method: ", userTypes)
			}
		}

		for _, userType := range userTypes {

			_, err := VerifyToken(token, userType)

			if err != nil {
				continue
			}

			valid = true
			break
		}

		if !valid {
			err := errors.New("Usuario Invalido")
			DenyPermision(ctx, err)
		}
	}
}

//DenyPermision =
func DenyPermision(ctx *context.Context, err error) {

	ctx.Output.SetStatus(401)
	ctx.Output.Header("Content-Type", "application/json")

	message := MessageResponse{
		Message:       "Permission Deny",
		PrettyMessage: "Permiso Denegado",
		Error:         err.Error(),
	}

	res, _ := json.Marshal(message)

	ctx.Output.Body([]byte(string(res)))
	return
}
