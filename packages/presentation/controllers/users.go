package controllers

import (
	"enube-challenge/packages/domains"
	"enube-challenge/packages/infra/dto"
	"enube-challenge/packages/infra/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usersUseCase domains.UsersUseCase
}

func NewUserController(usersUseCase domains.UsersUseCase) *UserController {
	return &UserController{
		usersUseCase: usersUseCase,
	}
}

// Create godoc
// @Summary     Create a new user
// @Description Creates a new user in the system. The user data is provided in the request body.
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       user body dto.CreateUserRequestDTO true "Create User Request"
// @Success     201  {object} domains.User
// @Failure     400  {string} "Invalid request body"
// @Failure     409  {string} "User already exists"
// @Router      /api/v1/users [post]
func (uc *UserController) Create(ctx *gin.Context) {
	var req dto.CreateUserRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid request body",
		})
		return
	}

	user := domains.User{
		Email:    req.Email,
		Username: req.Name,
	}

	createdUser, err := uc.usersUseCase.Create(ctx.Request.Context(), &user)
	if err != nil {
		errors.UserAlreadyExistHandler(ctx, ctx.Error(err))
		return
	}

	response := domains.HttpResponse{
		Code:    http.StatusCreated,
		Body:    createdUser,
		Message: "User created successfully",
	}
	ctx.JSON(http.StatusCreated, response)
}

// FindByEmail godoc
// @Summary     Find a user in the database by email
// @Description Retrieves a user from the database based on the provided email address.
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       email path string true "User Email"
// @Success     200  {object} domains.User
// @Failure     400  {string} "Bad request"
// @Failure     404  {string} "User not found"
// @Router      /api/v1/users/{email} [get]
func (uc *UserController) FindByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	u, err := uc.usersUseCase.FindByEmail(ctx.Request.Context(), email)
	if err != nil {
		errors.UserNotFoundHandler(ctx, ctx.Error(err))
		return
	}

	response := domains.HttpResponse{
		Code:    http.StatusOK,
		Message: "User successfully found",
		Body:    u,
	}
	ctx.JSON(http.StatusOK, response)
}
