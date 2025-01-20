// migrations/migrate.go - alternative approach
package migrations

import (
	"log"

	"github.com/deepanshumishraa/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	log.Println("Running migrations...")

	// Drop the table if it exists
	err := db.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatal("Failed to drop table:", err)
	}

	// Run auto migrations to create the table with new structure
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Migrations completed successfully")
}
