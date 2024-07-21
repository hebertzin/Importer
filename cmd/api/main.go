package main

import (
	_ "enube-challenge/docs"
	"enube-challenge/packages/config"
	"enube-challenge/packages/database"
	"enube-challenge/packages/logging"
	"enube-challenge/packages/routes"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Enube challenge
// @version 1.0
// @description Enube challenge

// @contact.name Hebert santos
// @contact.url https://www.hebertzin.com/
// @contact.email hebertsantosdeveloper@gmail.com

// @BasePath /api/v1
func main() {
	logging.InitLogger()
	c := config.LoadConfig()
	db := database.ConnectDatabase(c)

	err := database.Migrate(db)
	if err != nil {
		panic("Error migrating database: " + err.Error())
	}

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRouter(r, db)
	routes.AuthRouter(r, db)
	routes.Importer(r, db)

	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
