package initializers

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/lordofthemind/gormGo/types"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectToPostgresql() error{
// 	var err error
// 	postgresqlConnectionString := os.Getenv("PG_DB_CONNECTION")
// 	DB, err = gorm.Open(postgres.Open(postgresqlConnectionString), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connected to database")
// 	return nil
// }

// func SyncPostgresql() error {
// 	err := DB.AutoMigrate(&types.PersonType{})
// 	if err != nil {
// 		return fmt.Errorf("failed to auto migrate PersonType model: %w", err)
// 	}

// 	log.Println("Synchronized the PersonType model")
// 	log.Println("Synchronized the database")

// 	return nil
// }
