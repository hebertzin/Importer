package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": "This route upload xlsx file",
	})
}
