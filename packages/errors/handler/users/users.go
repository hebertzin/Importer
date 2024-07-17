package users

import (
	errors "enube-challenge/packages/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAlreadyExistHandler(ctx *gin.Context, err error) {
	statusCode := http.StatusNotFound
	message := "User already exist"

	if err == errors.ErrUserAlreadyExist {
		statusCode = http.StatusConflict
		message = "User already exist"
	}

	ctx.JSON(statusCode, gin.H{
		"msg":  message,
		"code": statusCode,
	})
}
