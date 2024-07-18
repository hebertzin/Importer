package controllers

import (
	"enube-challenge/packages/dto"
	"enube-challenge/packages/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctrl *AuthController) SignInHandler(c *gin.Context) {
	var req dto.LoginRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.authService.Auth(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
