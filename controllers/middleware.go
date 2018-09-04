package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
)

// MiddlewareController operations for Middleware
type MiddlewareController struct {
	BaseController
}

type loginToken struct {
	Type string `json:"tipo"`
	ID   int64  `json:"id"`
	jwt.StandardClaims
}

//Auth =
func Auth(c *context.Context) bool {
	success := VerifyToken(c.Input.Header("Authorization"))

	return success
}

//VerifyToken =
func VerifyToken(tokenString string) bool {

	hmacSampleSecret := []byte("bazam")

	token, err := jwt.ParseWithClaims(tokenString, &loginToken{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	claims, ok := token.Claims.(*loginToken)

	if !ok && !token.Valid {
		fmt.Println(err)
		return false
	}

	fmt.Printf("%v %v", claims.Type, claims.StandardClaims.ExpiresAt)
	return true
}

// GenerateToken =
func (c *BaseController) GenerateToken(userType string, id int64) (token string) {

	hmacSampleSecret := []byte("bazam")

	now := time.Now()

	// Create the Claims
	claims := loginToken{
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

	//fmt.Println("Token Generado")

	if err != nil {
		c.BadRequest(err)
	}

	//fmt.Printf("%v", ss)
	//fmt.Println(token)
	return token
}
