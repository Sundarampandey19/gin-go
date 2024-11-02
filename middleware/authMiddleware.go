package middleware

import (
	"net/http"
	"strings"
	"todo/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader =="" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error" : "Authorization header missing"})
			c.Abort()
			return 
		}

		tokenStr := authHeader[len("Bearer "):]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token)(interface{},error){
			return []byte(utils.JwtKey),nil 		
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims , ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId" , claims["userId"])
			c.Next()
		}else {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid token"})
			c.Abort()
			return
		}


	}
}