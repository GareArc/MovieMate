package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GareArc/MovieMate/internal/config"
	"github.com/GareArc/MovieMate/internal/db"
)

func initServer() {
	// Create a new server
	r := gin.Default()

	// Add a handler for the root page
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// TODO: Add middilewares

	// TODO: Add routes
	log.Printf("Server starting at %s:%d", config.MainConfig.String("server.host"), config.MainConfig.Int("server.port"))
	if err := r.Run(fmt.Sprintf("%s:%d", config.MainConfig.String("server.host"), config.MainConfig.Int("server.port"))); err != nil {
		log.Fatalf("error in server: %v", err)
	}
}

func Run() {
	config.InitConfig()
	db.InitDB()
	initServer()
}
