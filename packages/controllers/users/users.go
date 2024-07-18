package users

import (
	dto "enube-challenge/packages/controllers/dto/users"
	handle "enube-challenge/packages/errors/handler/users"
	models "enube-challenge/packages/models/users"
	s "enube-challenge/packages/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService s.UsersService
}

func NewUserController(s s.UsersService) *UserController {
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
		handle.UserAlreadyExistHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) FindByEmail(ctx *gin.Context) {
	var email string = ctx.Param("email")

	u, err := uc.userService.FindByEmail(ctx, email)

	if err != nil {
		handle.UserNotFoundHandler(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
