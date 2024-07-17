package api

import (
	"enube-challenge/packages/database"
	"enube-challenge/packages/database/migrations"
	"enube-challenge/packages/routes/users"

	"github.com/gin-gonic/gin"
)

func SetupServer() {
	db := database.ConnectDatabase()

	err := migrations.Migrate(db)
	if err != nil {
		panic("Erro ao executar migração: " + err.Error())
	}

	r := gin.Default()

	users.UserRouter(r, db)

	r.Run(":8080")
}