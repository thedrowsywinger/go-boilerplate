package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"go-boilerplate/controllers"
	"go-boilerplate/middlewares"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/welcome", middlewares.AuthenticateJWT(), controllers.Welcome)
}
