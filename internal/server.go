package internal

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/GareArc/MovieMate/internal/config"
	"github.com/GareArc/MovieMate/internal/db"
	"github.com/GareArc/MovieMate/internal/router"
)

func initServer() {
	config := config.GetStaticConfig()
	// Create a new server
	r := gin.Default()

	// TODO: Add global middilewares

	// TODO: Add routes
	router.Router(r)

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
