package users

import (
	controllers "enube-challenge/packages/controllers/users"
	repository "enube-challenge/packages/repository/users"
	services "enube-challenge/packages/services/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(router *gin.Engine, db *gorm.DB) {
	usersRepository := repository.NewUsersRepository(db)
	usersService := services.NewUsersService(usersRepository)
	userControllers := controllers.NewUserController(usersService)

	usersGroup := router.Group("/api/v1")
	{
		usersGroup.POST("/users", userControllers.Create)
	}
}
