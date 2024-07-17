package api

import (
	"enube-challenge/packages/database"
	"enube-challenge/packages/routes/users"

	"github.com/gin-gonic/gin"
)

func SetupServer() {
	db := database.ConnectDatabase()

	r := gin.Default()

	users.UserRouter(r, db)

	r.Run(":8080")
}
