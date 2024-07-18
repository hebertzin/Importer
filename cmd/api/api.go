package api

import (
	"enube-challenge/packages/database"
	"enube-challenge/packages/logging"
	"enube-challenge/packages/routes"
	"github.com/gin-gonic/gin"
)

func SetupServer() {
	logging.InitLogger()
	db := database.ConnectDatabase()

	err := database.Migrate(db)
	if err != nil {
		panic("Error migrating database: " + err.Error())
	}

	r := gin.Default()

	routes.UserRouter(r, db)
	routes.AuthRouter(r, db)

	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
