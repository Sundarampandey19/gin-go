package utils

import (
	"os"
	"log"
	"time"
    "github.com/joho/godotenv"
	"github.com/dgrijalva/jwt-go"
)


var JwtKey []byte

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    JwtKey = []byte(os.Getenv("JWT_SECRET"))
}


func GenerateToken(userId string) (string,error) {
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId" : userId,
		"exp" : time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(JwtKey) 
}