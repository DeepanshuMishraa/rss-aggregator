// migrations/migrate.go
package migrations

import (
	"log"

	"github.com/deepanshumishraa/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	log.Println("Running migrations...")

	// Add your models here
	err := db.AutoMigrate(
		&models.User{},
		// Add more models here
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Migrations completed successfully")
}
