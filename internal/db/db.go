package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/GareArc/MovieMate/internal/config"
)

var (
	MainDB *sql.DB
)

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s TimeZone=%s",
		config.MainConfig.String("db.host"),
		config.MainConfig.String("db.user"),
		config.MainConfig.String("db.password"),
		config.MainConfig.String("db.port"),
		config.MainConfig.String("db.ssl"),
		config.MainConfig.String("db.timezone"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	postgresDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// check if dbname exists
	rows, err := postgresDB.Query("SELECT 1 FROM pg_database WHERE datname = $1", config.MainConfig.String("db.dbname"))
	if err != nil {
		log.Fatal(err)
	}

	if !rows.Next() {
		// create dbname
		_, err := postgresDB.Exec(fmt.Sprintf("CREATE DATABASE %s", config.MainConfig.String("db.dbname")))
		if err != nil {
			log.Fatal(err)
		}
	}
	postgresDB.Close()

	// reconnect to dbname
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.MainConfig.String("db.host"),
		config.MainConfig.String("db.user"),
		config.MainConfig.String("db.password"),
		config.MainConfig.String("db.dbname"),
		config.MainConfig.String("db.port"),
		config.MainConfig.String("db.ssl"),
		config.MainConfig.String("db.timezone"))

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
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

	MainDB = postgresDB

	return nil
}
