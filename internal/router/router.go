package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	user_controller "github.com/GareArc/MovieMate/internal/controller"
	"github.com/GareArc/MovieMate/internal/middlewares"
)

func Router(r *gin.Engine) {
	healthCheckRouter(r)
	v1_api := VersionRouterHead(r, "v1")
	initUserRouter(v1_api)
}

func VersionRouterHead(r *gin.Engine, version string) *gin.RouterGroup {
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

	UserGroup.POST("/register", user_controller.RegisterUser)
	UserGroup.POST("/login", user_controller.LoginUser)
	UserGroup.GET("/me", middlewares.RequireLogin(), user_controller.CheckActiveUser)

}
