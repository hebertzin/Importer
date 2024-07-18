package controllers

import (
	"net/http"

	"enube-challenge/packages/services/authentication"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService authentication.AuthService
}

func NewAuthController(authService authentication.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctrl *AuthController) SignInHandler(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.authService.Auth(c.Request.Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
