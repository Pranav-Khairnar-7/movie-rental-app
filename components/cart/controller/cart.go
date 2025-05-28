package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-rental-app/components/cart/service"
	custError "github.com/movie-rental-app/errors"
	model "github.com/movie-rental-app/models"
	uuid "github.com/satori/go.uuid"
)

type CartController struct {
	CartService *service.CartService
}

func NewCartController(svc *service.CartService) *CartController {
	return &CartController{CartService: svc}
}

func (c *CartController) CreateCart(ctx *gin.Context) {
	var cart model.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.CartService.CreateCart(ctx, &cart); err != nil {
		var httpErr *custError.HTTPError
		if errors.As(err, &httpErr) {
			ctx.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, cart)
}

func (c *CartController) GetCartByID(ctx *gin.Context) {
	var (
		cartID uuid.UUID
	)

	if l := ctx.Param("cartID"); l != "" {
		cartID = uuid.FromStringOrNil(l)
	}

	var cart model.Cart

	err := c.CartService.GetCartByID(ctx, &cart, cartID)
	if err != nil {
		var httpErr *custError.HTTPError
		if errors.As(err, &httpErr) {
			ctx.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, cart)
}
