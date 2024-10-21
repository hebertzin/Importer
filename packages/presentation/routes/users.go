package routes

import (
	"enube-challenge/packages/infra/db/repository"
	"enube-challenge/packages/presentation/controllers"
	"enube-challenge/packages/presentation/middlewares"
	"enube-challenge/packages/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(router *gin.Engine, db *gorm.DB) {
	usersRepository := repository.NewUsersRepository(db)
	usersUseCase := usecases.NewUserUseCase(usersRepository)
	userControllers := controllers.NewUserController(usersUseCase)
	jwtService := usecases.NewJWTUseCase()
	usersGroup := router.Group("/api/v1")
	{
		usersGroup.POST("/users", userControllers.Create)
		usersGroup.GET("/users/:email", middleware.AuthMiddleware(jwtService), userControllers.FindByEmail)
	}
}
