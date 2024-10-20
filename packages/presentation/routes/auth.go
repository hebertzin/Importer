package routes

import (
	"enube-challenge/packages/infra/db/repository"
	"enube-challenge/packages/presentation/controllers"
	"enube-challenge/packages/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(router *gin.Engine, db *gorm.DB) {
	jwtService := usecases.NewJWTUseCase()
	userRepository := repository.NewUsersRepository(db)
	authService := usecases.NewAuthUseCase(userRepository, jwtService)
	authController := controllers.NewAuthController(authService)
	authGroup := router.Group("/api/v1")
	{
		authGroup.POST("/authentication/sign-in", authController.SignInHandler)
	}
}
