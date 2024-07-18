package controllers

import (
	"enube-challenge/packages/dto"
	"enube-challenge/packages/errors"
	"enube-challenge/packages/models"
	"enube-challenge/packages/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UsersService
}

func NewUserController(s services.UsersService) *UserController {
	return &UserController{
		userService: s,
	}
}

func (uc *UserController) Create(ctx *gin.Context) {
	var req dto.CreateUserRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid request body",
		})
		return
	}

	user := models.Users{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Name,
	}

	createdUser, err := uc.userService.Create(ctx.Request.Context(), &user)
	if err != nil {
		errors.UserAlreadyExistHandler(ctx, ctx.Error(err))
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) FindByEmail(ctx *gin.Context) {
	var email string = ctx.Param("email")

	u, err := uc.userService.FindByEmail(ctx, email)

	if err != nil {
		errors.UserNotFoundHandler(ctx, ctx.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
