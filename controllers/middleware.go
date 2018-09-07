package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

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

	fmt.Println(tokenString, userType)
	fmt.Println(claims)

	//fmt.Printf("%v %v", claims.Type, claims.StandardClaims.ExpiresAt)

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

//GetValidation =
func GetValidation(route string) (validation map[string]map[string][]string) {

	portfolios := map[string]map[string][]string{
		"POST":   {"/": {"Admin"}},
		"GET":    {"portfolios/:id": {"Admin", "Client"}},
		"PUT":    {"/:id": {"Admin"}},
		"DELETE": {"/:id": {"Admin"}},
	}

	validations := map[string]map[string]map[string][]string{
		"portfolios": portfolios,
	}

	return validations[route]
}

//Middleware =
func Middleware(route string) func(ctx *context.Context) bool {

	validation := GetValidation(route)

	return func(ctx *context.Context) bool {

		token := ctx.Input.Header("Authorization")
		method := validation[ctx.Input.Method()]

		//beego.Debug("Token: ", token)
		//beego.Debug("Method: ", method)

		for url, usersTypes := range method {

			//beego.Debug("Url: ", ctx.Input.URL(), "Match: ", url)

			urlFound := strings.HasSuffix(ctx.Input.URL(), url)

			//beego.Debug("Tiene El sufijo: ", urlFound)

			if !urlFound {
				return false
			}

			for _, userType := range usersTypes {

				_, err := VerifyToken(token, userType)

				if err != nil {
					DenyPermision(ctx, err)
				}
			}
		}

		return false
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
