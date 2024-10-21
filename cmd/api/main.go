package main

import (
	_ "enube-challenge/docs"
	"enube-challenge/packages/infra/config"
	database2 "enube-challenge/packages/infra/db"
	"enube-challenge/packages/infra/logging"
	appRoutes "enube-challenge/packages/presentation/routes"
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
	dbConfig := config.LoadConfig()
	db := database2.ConnectDatabase(dbConfig)

	err := database2.Migrate(db)
	if err != nil {
		logging.Log.Fatal("Error migrating db", zap.Error(err))
	}

	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	appRoutes.UserRouter(router, db)
	appRoutes.AuthRouter(router, db)
	appRoutes.Suppliers(router, db)

	if err := router.Run(":8080"); err != nil {
		logging.Log.Fatal("Failed to start server", zap.Error(err))
	}
}
