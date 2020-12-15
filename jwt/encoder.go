package jwt

import (
	"authentication/model"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)


func GenerateToken(user string, pass string,minutes string) (string, error){
	var claims model.MyClaims
	var jwtKey = []byte("my-secret-key")

	expTime,err := strconv.Atoi(minutes)

	tokenTime := time.Duration(expTime) * time.Minute


	if err != nil{
		return "",err
	}
	expiration := time.Now().Add(tokenTime)

	claims = model.MyClaims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err !=nil{
		return "",err
	}
	return tokenString,nil
}