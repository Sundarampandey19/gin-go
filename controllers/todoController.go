package controllers

import (
	"context"
	"net/http"
	database "todo/Database"
	"todo/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateTodo(c *gin.Context){
	userId, _ := c.Get("userId")
	var todo models.Todo
	
	if err:= c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid Request"})
		return
	}

	todo.UsedID , _  = primitive.ObjectIDFromHex(userId.(string))
	collection := database.Client.Database("todoapp").Collection("todos")
	result , err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failed to create todo"})
		return 
	}	
	c.JSON(http.StatusCreated, gin.H{"todo_id" : result.InsertedID})
}

func GetTodos(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{"todo_id" :"Returns all the todos"})


}

func UpdateTodo(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{"todo_id" :"request to update todos"})


}

func DeleteTodo(c *gin.Context){
	c.JSON(http.StatusAccepted, gin.H{"todo_id" :"Request to delete a todo"})


}