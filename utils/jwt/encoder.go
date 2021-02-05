package jwt

import (
	"authentication/entities"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

//Method to generate a jwt token using a secret key and 30 minutes of lifespan
//Returns a token string
func GenerateToken(id int ,ip string,minutes string,key string) (string, error){
	var claims entities.MyClaims
	var jwtKey = []byte(key)

	expTime,err := strconv.Atoi(minutes)

	if err != nil{
		return "",err
	}

	tokenTime := time.Duration(expTime) * time.Minute
	expiration := time.Now().Add(tokenTime)

	claims = entities.MyClaims{
		Ip: ip,
		Id: id,
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