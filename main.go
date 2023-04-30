package main

import (
	"github.com/Karthika-Rajagopal/jwt-go/controllers"
	"github.com/Karthika-Rajagopal/jwt-go/initializers"
	"github.com/Karthika-Rajagopal/jwt-go/middleware"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
func main(){
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run() 
}