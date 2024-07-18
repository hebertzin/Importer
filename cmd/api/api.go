package api

import (
	logger "enube-challenge/packages/config/logging"
	"enube-challenge/packages/database"
	"enube-challenge/packages/database/migrations"
	"enube-challenge/packages/routes/auth"
	"enube-challenge/packages/routes/users"

	"github.com/gin-gonic/gin"
)

func SetupServer() {
	logger.InitLogger()
	db := database.ConnectDatabase()

	err := migrations.Migrate(db)
	if err != nil {
		panic("Error migrating database: " + err.Error())
	}

	r := gin.Default()

	users.UserRouter(r, db)
	auth.AuthRouter(r, db)

	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
