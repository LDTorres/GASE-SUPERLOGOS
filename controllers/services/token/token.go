package token

import (
	"errors"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

//JwtToken =
type JwtToken struct {
	ID int 
	Type string
	jwt.StandardClaims
}

var secretWord = []byte("bazam")

func GenerateTimeToken (ID int, tokenType string , timeArgs ...int) (token string, err error) {

	now := time.Now()

	timeValues := []int{14, 0, 0}

	for key, timeArg := range timeArgs {
		timeValues[key] = timeArg
	}

	// Create the Claims
	claims := JwtToken{
		ID,
		tokenType,
		jwt.StandardClaims{
			ExpiresAt: now.AddDate(timeValues[2], timeValues[1], timeValues[0]).Unix(),
			Issuer:    "test",
		},
	}

	var newToken *jwt.Token
	newToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = newToken.SignedString(secretWord)

	return
	
}

func ValidateTimeToken (tokenString string, tokenType string) (decodedToken *JwtToken, err error){

	if tokenString == "" {
		err = errors.New("Empty Token")
		return 
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtToken{}, func(token *jwt.Token) (interface{}, error) {
		return secretWord, nil
	})

	if err != nil {
		return 
	}

	claims, ok := token.Claims.(*JwtToken)

	if !ok || !token.Valid {
		return 
	}

	
	if claims.Type != tokenType {
		return nil, errors.New("Invalid Token's Type")
	}

	return claims, nil

}