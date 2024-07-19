package controllers

import (
	"enube-challenge/packages/domain"
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

// SignInHandler godoc
// @Summary Sign-in a user
// @Description This function logs in a user and generates a token.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body dto.LoginRequestDTO true "Login Request"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/v1/authentication/sign-in [post]
func (ctrl *AuthController) SignInHandler(c *gin.Context) {
	var req dto.LoginRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response := domain.HttpResponse{
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
			Body:    "",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	authResponse, err := ctrl.authService.Auth(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		response := domain.HttpResponse{
			Message: authResponse.Message,
			Code:    authResponse.Code,
			Body:    authResponse.Body,
		}
		c.JSON(authResponse.Code, response)
		return
	}

	response := domain.HttpResponse{
		Message: "User logged in successfully",
		Code:    http.StatusOK,
		Body:    authResponse.Body,
	}

	c.JSON(http.StatusOK, response)
}
