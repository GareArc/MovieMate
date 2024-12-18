package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/GareArc/MovieMate/internal/controller"
	"github.com/GareArc/MovieMate/internal/middlewares"
)

func Router(r *gin.Engine) {
	healthCheckRouter(r)
	v1_api := versionRouterHead(r, "v1")
	initUserRouter(v1_api)
	initMovieRouter(v1_api)
}

func versionRouterHead(r *gin.Engine, version string) *gin.RouterGroup {
	return r.Group(fmt.Sprintf("/api/%s", version))
}

func healthCheckRouter(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I am alive",
		})
	})
}

func initUserRouter(r *gin.RouterGroup) {
	UserGroup := r.Group("/user")
	user_controller := controller.UserController{}

	UserGroup.POST("/register", user_controller.RegisterUser)
	UserGroup.POST("/login", user_controller.LoginUser)
	UserGroup.GET("/me", middlewares.RequireLogin(), user_controller.CheckActiveUser)

}

func initMovieRouter(r *gin.RouterGroup) {
	MovieGroup := r.Group("/movie")
	movie_controller := controller.MovieController{}

	MovieGroup.GET("/", movie_controller.GetMovieInfo)
	MovieGroup.GET("/showtime", movie_controller.GetShowTimes)
}
