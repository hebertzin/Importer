package main

import (
	"enube-challenge/packages/config"
	"enube-challenge/packages/database"
	"enube-challenge/packages/logging"
	"enube-challenge/packages/routes"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
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
	defer logging.Log.Sync()
	c := config.LoadConfig()
	db := database.ConnectDatabase(c)

	err := database.Migrate(db)
	if err != nil {
		logging.Log.Fatal("Error migrating database", zap.Error(err))
	}

	r := gin.Default()

	r.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRouter(r, db)
	routes.AuthRouter(r, db)
	routes.Suppliers(r, db)

	if err := r.Run(":8080"); err != nil {
		logging.Log.Fatal("Failed to start server", zap.Error(err))
	}
}
