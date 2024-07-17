package users

import (
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
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Some error has been ocurred",
		})
	}

	u, err := uc.userService.Create(ctx, &user)

	if err != nil {
		panic("Some error has been ocurred")
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"user": u,
	})
}

func (uc *UserController) FindByEmail(ctx *gin.Context) {
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Some error has been ocurred",
		})
	}

	u, _ := uc.userService.Create(ctx, &user)

	ctx.JSON(http.StatusCreated, gin.H{
		"user": u,
	})
}
