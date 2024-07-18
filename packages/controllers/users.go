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

// Create @Summary Add a new user
// @Description This handler func create a new user in database
// @Tags users
// @Accept  json
// @Produce  json
// @Param product body dto.CreateUserRequestDTO
// @Success 201
// @Router /api/v1/users [post]
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

// FindByEmail @Summary Get a new user
// @Description This handler func get a user by email
// @Tags users
// @Accept  json
// @Produce  json
// @Param product params email
// @Success 200
// @Router /api/v1/users/:email [get]
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
