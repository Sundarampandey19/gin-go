package controllers

import (
	"context"
	"fmt"
	"net/http"
	"todo/Database"
	"todo/models"
	"todo/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	fmt.Println(c)

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	collection := database.Client.Database("todoapp").Collection("users")
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user_id": result.InsertedID})

}


func Login (c * gin.Context){
	var input struct {
		Username	 string `json:"username"`
		Password 	 string `json:"password"`
	}

	if err := c.BindJSON(&input) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid input"})
		return 
	}

	collection := database.Client.Database("todoapp").Collection("users")

	var user models.User 
	if err := collection.FindOne(context.Background(), bson.M{"username" : input.Username}).Decode(&user); err !=nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid username or password"})
		return 
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password) , []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid username or password"})
		return 
	}

	token, _ := utils.GenerateToken(user.ID.Hex())
	c.JSON(http.StatusOK, gin.H{"token" : token})




}