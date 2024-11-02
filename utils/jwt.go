package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)


var JwtKey = []byte("Secretkey")

func GenerateToken(userId string) (string,error) {
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId" : userId,
		"exp" : time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(JwtKey) 
}