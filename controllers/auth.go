package controllers

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// AuthController operations for Middleware
type AuthController struct {
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
