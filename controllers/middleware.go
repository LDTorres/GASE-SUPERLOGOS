package controllers

import (
	"encoding/json"
	"errors"
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
	ID   int64  `json:"id"`
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

	//fmt.Printf("%v %v", claims.Type, claims.StandardClaims.ExpiresAt)

	if claims.Type != userType {
		return nil, errors.New("Invalid User")
	}

	return claims, nil
}

// GenerateToken =
func (c *BaseController) GenerateToken(userType string, id int64) (token string) {

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

//GlobalMiddleware =
var GlobalMiddleware = func(ctx *context.Context) {

	ValidateUrls := map[string][]string{
		// Clients
		"/clients": {"GET", "Admin"},
		// Activities
		"/activities": {"GET", "Admin"},
		//
	}

	for url, options := range ValidateUrls {

		userType := options[1]

		if strings.HasSuffix(ctx.Input.URL(), url) && ctx.Input.Method() == options[0] {
			_, err := VerifyToken(ctx.Input.Header("Authorization"), userType)

			if err != nil {
				DenyPermision(ctx, err)
			}

			break
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
