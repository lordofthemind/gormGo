package tests

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lordofthemind/gormGo/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvVariablesForTesting() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	log.Println("Environment variables loaded")
	fmt.Println("Environment variables loaded")
	return nil
}

var MOCK_DB *gorm.DB

func ConnectToPostgresqlForTesting() error {
	postgresqlConnectionStringForTesting := os.Getenv("PG_DB_CONNECTION_FOR_TESTING")
	db, err := gorm.Open(postgres.Open(postgresqlConnectionStringForTesting), &gorm.Config{})
	if err != nil {
		return err
	}

	MOCK_DB = db

	log.Println("Connected to test database")
	fmt.Println("Connected to test database")
	return nil
}

func SyncPostgresqlForTesting() error {
	err := MOCK_DB.AutoMigrate(&types.PersonType{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate PersonType model: %w", err)
	}

	log.Println("Synchronized the PersonType model")
	log.Println("Synchronized the database")
	fmt.Println("Synchronized the PersonType model")
	fmt.Println("Synchronized the database")

	return nil
}

func InitializeForTesting() error {
	err := LoadEnvVariablesForTesting()
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	err = ConnectToPostgresqlForTesting()
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	err = SyncPostgresqlForTesting()
	if err != nil {
		return fmt.Errorf("failed to sync PostgreSQL: %w", err)
	}

	fmt.Println("Initialization completed successfully")
	return nil
}
