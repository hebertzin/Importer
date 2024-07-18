package users

import (
	controllers "enube-challenge/packages/controllers/users"
	middleware "enube-challenge/packages/middlewares"
	repository "enube-challenge/packages/repository/users"
	"enube-challenge/packages/services/jwt"
	services "enube-challenge/packages/services/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(router *gin.Engine, db *gorm.DB) {
	usersRepository := repository.NewUsersRepository(db)
	usersService := services.NewUsersService(usersRepository)
	userControllers := controllers.NewUserController(usersService)

	jwtService := jwt.NewJWTService()

	usersGroup := router.Group("/api/v1")
	{
		usersGroup.POST("/users", userControllers.Create)
		usersGroup.GET("/users/:email", middleware.AuthMiddleware(jwtService), userControllers.FindByEmail)
	}
}
