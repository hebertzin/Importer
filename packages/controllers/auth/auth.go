package controllers

import (
	"net/http"

	dto "enube-challenge/packages/controllers/dto/auth"
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
