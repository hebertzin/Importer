package controllers

import (
	"enube-challenge/packages/domain"
	"enube-challenge/packages/dto"
	"enube-challenge/packages/errors"
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

// Create godoc
// @Summary     Create a new user
// @Description  Creates a new user in the system. The user data is provided in the request body.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body dto.CreateUserRequestDTO true "Create User Request"
// @Success      201  {object} models.Users
// @Failure      400  Invalid request body
// @Failure      409  User already exist
// @Router       /api/v1/users [post]
func (uc *UserController) Create(ctx *gin.Context) {
	var req dto.CreateUserRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid request body",
		})
		return
	}

	user := domain.User{
		Email:    req.Email,
		Username: req.Name,
	}

	createdUser, err := uc.userService.Create(ctx.Request.Context(), &user)
	if err != nil {
		errors.UserAlreadyExistHandler(ctx, ctx.Error(err))
		return
	}

	response := domain.HttpResponse{
		Code:    http.StatusCreated,
		Body:    createdUser,
		Message: "User created successfully",
	}
	ctx.JSON(http.StatusOK, response)
}

// FindByEmail godoc
// @Summary     Find a user in the database by email
// @Description  Retrieves a user from the database based on the provided email address.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        email path string true "User Email"
// @Success      200
// @Failure      400  {string} Bad request
// @Failure      404  {string} Not found
// @Router       /api/v1/users/{email} [get]
func (uc *UserController) FindByEmail(ctx *gin.Context) {
	var email string = ctx.Param("email")

	u, err := uc.userService.FindByEmail(ctx, email)

	if err != nil {
		errors.UserNotFoundHandler(ctx, ctx.Error(err))
		return
	}

	response := domain.HttpResponse{
		Code:    http.StatusOK,
		Message: "User successfully find",
		Body:    u,
	}
	ctx.JSON(http.StatusOK, response)
}
