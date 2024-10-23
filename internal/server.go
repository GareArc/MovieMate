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
	config := config.GetStaticConfig()
	// Create a new server
	r := gin.Default()

	// Add a handler for the root page
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// TODO: Add middilewares

	// TODO: Add routes
	log.Printf("Server starting at %s:%d", config.String("server.host"), config.Int("server.port"))
	if err := r.Run(fmt.Sprintf("%s:%d", config.String("server.host"), config.Int("server.port"))); err != nil {
		log.Fatalf("error in server: %v", err)
	}
}

func initDatabase() {
	db.InitDB()
	enable_migrate := config.GetStaticConfig().Bool("db.migrate")
	if enable_migrate {
		err := db.Migrate()
		if err != nil {
			log.Fatalf("error in migratation: %v", err)
		}
	}
}

func Run() {
	config.InitConfig()
	initDatabase()
	initServer()
}
