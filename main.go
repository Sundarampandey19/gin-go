package main

import (
    "todo/controllers"
    "todo/middleware"
    "todo/Database"
    "github.com/gin-gonic/gin"
)




func main() {


    database.ConnectMongoDB()
    router := gin.Default()

    router.POST("/register" , controllers.Register)
    router.POST("/login" , controllers.Login)

    authozised :=router.Group("/")

    authozised.Use(middleware.AuthMiddleware())
    {
        authozised.GET("/todos", controllers.GetTodos)
        authozised.POST("/todos", controllers.CreateTodo)
        authozised.PUT("/todos/:id", controllers.UpdateTodo)
        authozised.DELETE("/todos/:id", controllers.DeleteTodo)

    }

    router.Run(":8080")


	
}