package routes

import (
	"enube-challenge/packages/presentation/controllers"
	"enube-challenge/packages/presentation/middlewares"
	"enube-challenge/packages/repository"
	"enube-challenge/packages/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(router *gin.Engine, db *gorm.DB) {
	usersRepository := repository.NewUsersRepository(db)
	usersService := services.NewUsersService(usersRepository)
	userControllers := controllers.NewUserController(usersService)

	jwtService := services.NewJWTService()

	usersGroup := router.Group("/api/v1")
	{
		usersGroup.POST("/users", userControllers.Create)
		usersGroup.GET("/users/:email", middleware.middleware.AuthMiddleware(jwtService), userControllers.FindByEmail)
	}
}
