package database

import (
	"air-trail-backend/utils/env"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Pgsql *gorm.DB

func init() {
	var (
		DB_HOST     = env.GetEnv("POSTGRESQL_HOST", "localhost")
		DB_USER     = env.GetEnv("POSTGRESQL_USERNAME", "postgres")
		DB_PASSWORD = env.GetEnv("POSTGRESQL_PASSWORD", "")
		DB_NAME     = env.GetEnv("POSTGRESQL_DBNAME", "air_trail_db")
		DB_PORT     = env.GetEnv("POSTGRESQL_PORT", "5432")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	Pgsql = db
	log.Printf("Connected to database %s at %s:%s", DB_NAME, DB_HOST, DB_PORT)
}
