package routes

import (
	"enube-challenge/packages/presentation/controllers"
	"enube-challenge/packages/repository"
	"enube-challenge/packages/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(router *gin.Engine, db *gorm.DB) {

	jwtService := services.NewJWTService()

	userRepository := repository.NewUsersRepository(db)

	authService := services.NewAuthService(userRepository, jwtService)

	authController := controllers.NewAuthController(authService)

	authGroup := router.Group("/api/v1")
	{
		authGroup.POST("/authentication/sign-in", authController.SignInHandler)
	}
}
