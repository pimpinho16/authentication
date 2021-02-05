package entities

import "github.com/dgrijalva/jwt-go"

type MyClaims struct{
	Id int
	Ip string
	jwt.StandardClaims
}
