package main

import (
	"jwt/controllers"
	"jwt/middlewares"
	"jwt/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()
	r := gin.Default()
	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.POST("/verify", controllers.VerifyOtp)
	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	r.Run((":8080"))
}
