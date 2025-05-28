package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-rental-app/components/user/service"
	custError "github.com/movie-rental-app/errors"
	model "github.com/movie-rental-app/models"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{UserService: svc}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserService.CreateUser(ctx.Request.Context(), &user); err != nil {
		var httpErr *custError.HTTPError
		if errors.As(err, &httpErr) {
			ctx.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
