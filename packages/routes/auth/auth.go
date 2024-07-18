package auth

import (
	controllers "enube-challenge/packages/controllers/auth"
	"enube-challenge/packages/repository/users"
	"enube-challenge/packages/services/authentication"
	"enube-challenge/packages/services/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(router *gin.Engine, db *gorm.DB) {

	jwtService := jwt.NewJWTService()

	userRepository := users.NewUsersRepository(db)

	authService := authentication.NewAuthService(userRepository, jwtService)

	authController := controllers.NewAuthController(authService)

	authGroup := router.Group("/api/v1")
	{
		authGroup.POST("/authentication/sign-in", authController.SignInHandler)
	}
}
