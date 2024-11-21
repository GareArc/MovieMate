package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/GareArc/MovieMate/internal/config"
	"github.com/GareArc/MovieMate/internal/types/model"
)

var (
	MainDB *gorm.DB
)

func InitDB() error {
	config := config.GetStaticConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s TimeZone=%s",
		config.String("db.host"),
		config.String("db.user"),
		config.String("db.password"),
		config.String("db.port"),
		config.String("db.ssl"),
		config.String("db.timezone"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	postgresDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// check if dbname exists
	rows, err := postgresDB.Query("SELECT 1 FROM pg_database WHERE datname = $1", config.String("db.dbname"))
	if err != nil {
		log.Fatal(err)
	}

	if !rows.Next() {
		// create dbname
		_, err := postgresDB.Exec(fmt.Sprintf("CREATE DATABASE %s", config.String("db.dbname")))
		if err != nil {
			log.Fatal(err)
		}
	}
	postgresDB.Close()

	// reconnect to dbname
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.String("db.host"),
		config.String("db.user"),
		config.String("db.password"),
		config.String("db.dbname"),
		config.String("db.port"),
		config.String("db.ssl"),
		config.String("db.timezone"))

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal(err)
	}

	postgresDB, err = db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// check if uuid-ossp exists
	rows, err = postgresDB.Query("SELECT 1 FROM pg_extension WHERE extname = 'uuid-ossp'")
	if err != nil {
		log.Fatal(err)
	}

	if !rows.Next() {
		// create uuid-ossp
		_, err := postgresDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
		if err != nil {
			log.Fatal(err)
		}
	}

	postgresDB.SetMaxIdleConns(10)
	postgresDB.SetMaxOpenConns(100)
	postgresDB.SetConnMaxLifetime(time.Hour)

	MainDB = db

	return nil
}

func Migrate() error {
	// Migrate the schemas
	MainDB.AutoMigrate(
		&model.User{},
		&model.Movie{},
		&model.Theater{},
		&model.Seat{},
		&model.MovieSchedule{},
		&model.Ticket{},
		&model.TicketPurchase{},
	)

	return nil
}
