package middleware

import (
	"enube-challenge/packages/domains"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyXLSXMiddleware(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.Next()
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, domains.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}
	defer file.Close()

	filename := header.Filename
	if !strings.HasSuffix(strings.ToLower(filename), ".xlsx") {
		c.JSON(http.StatusBadRequest, domains.HttpResponse{
			Message: "Invalid type, please upload .xlsx as an Excel file",
			Code:    http.StatusUnsupportedMediaType,
		})
		c.Abort()
		return
	}

	c.Next()
}
