package internal

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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

	log.Info().Msgf("Server starting at %s:%d", config.String("server.host"), config.Int("server.port"))
	if err := r.Run(fmt.Sprintf("%s:%d", config.String("server.host"), config.Int("server.port"))); err != nil {
		log.Fatal().Err(err).Msg("error in server")
	}
}

func initDatabase() {
	db.InitDB()
	enable_migrate := config.GetStaticConfig().Bool("db.migrate")
	if enable_migrate {
		err := db.Migrate()
		if err != nil {
			log.Fatal().Err(err).Msg("error in migratation")
		}
	}
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

}

func Run() {
	config.InitConfig()
	initLogger()
	initDatabase()
	initServer()
}
