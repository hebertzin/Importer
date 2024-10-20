package middleware

import (
	"enube-challenge/packages/domains"
	"enube-challenge/packages/usecases"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *usecases.JWTUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, domains.HttpResponse{
				Code:    http.StatusUnauthorized,
				Message: "Authorization header is required",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, domains.HttpResponse{
				Code:    http.StatusUnauthorized,
				Message: "Bearer token is required",
			})
			c.Abort()
			return
		}

		claims, err := jwtService.Verify(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domains.HttpResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
