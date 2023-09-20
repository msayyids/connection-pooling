package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var dbPool *pgxpool.Pool

func InitDb() *pgxpool.Pool {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	dbPool = pool // Assign the pool to the global variable
	return dbPool
}
