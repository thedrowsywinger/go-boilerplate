package routes

import (
	"go-boilerplate/controllers"
	"go-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/refresh-token", controllers.RefreshToken)
	r.GET("/welcome", middlewares.AuthenticateJWT(), controllers.Welcome)
}
