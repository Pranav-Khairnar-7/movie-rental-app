package router

import (
	"github.com/gin-gonic/gin"
	CartCon "github.com/movie-rental-app/components/cart/controller"
	MovieCon "github.com/movie-rental-app/components/movie/controller"
	UserCon "github.com/movie-rental-app/components/user/controller"
)

func GetRouter() *gin.Engine {

	router := gin.Default()
	router.GET("/health", healthCheck)
	return router
}

func healthCheck(c *gin.Context) {
	// This function will handle the health check route.
	// It will return a simple message indicating that the service is running.
	c.JSON(200, gin.H{
		"message": "Service is running",
	})
}

func RegisterUserRoutes(r *gin.Engine, userController *UserCon.UserController) {

	userGroup := r.Group("/users")
	{
		userGroup.POST("", userController.CreateUser)
	}

}

func RegisterMovieRoutes(r *gin.Engine, movieController *MovieCon.MovieController) {

	movieGroup := r.Group("/movies")
	{
		movieGroup.GET("", movieController.GetMovies)
		movieGroup.GET("/:movieID", movieController.GetMovieByID)

	}

}

func RegisterCartRoutes(r *gin.Engine, cartController *CartCon.CartController) {

	cartGroup := r.Group("/cart")
	{
		cartGroup.POST("", cartController.CreateCart)
		cartGroup.GET("/:cartID", cartController.GetCartByID)
	}

}
