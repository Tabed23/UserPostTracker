package controller

import (
	"net/http"

	"github.com/Tabed23/UserPostTracker/app/service"
	"github.com/Tabed23/UserPostTracker/app/types"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
func (c *UserController) CreateUser(ctx *gin.Context) {
	usr := types.User{}

	if err := ctx.ShouldBindJSON(&usr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	usrID, err := c.userService.CreateUser(ctx, &usr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user_id": usrID})
}
