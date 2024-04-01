package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/lordofthemind/gormGo/types"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	log.Println("Environment variables loaded")
	return nil
}

var DB *gorm.DB

func ConnectToPostgresql() error {
	postgresqlConnectionString := os.Getenv("PG_DB_CONNECTION")
	db, err := gorm.Open(postgres.Open(postgresqlConnectionString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	DB = db

	log.Println("Connected to database")
	return nil
}

func SyncPostgresql() error {
	err := DB.AutoMigrate(&types.PersonType{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate PersonType model: %w", err)
	}

	log.Println("Synchronized the PersonType model")
	log.Println("Synchronized the database")

	return nil
}

func Initialize() error {
	err := LoadEnvVariables()
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	err = ConnectToPostgresql()
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	err = SyncPostgresql()
	if err != nil {
		return fmt.Errorf("failed to sync PostgreSQL: %w", err)
	}

	log.Println("Initialization completed successfully")
	return nil
}
