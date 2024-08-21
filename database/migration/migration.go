package migration

import (
	"e-wallet/database"
	"e-wallet/models"

	"fmt"
	"log"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.Customer{},
		&models.WalletTransaction{})
	if err != nil {
		log.Fatal("Failed to migrate")
	}

	fmt.Println("Migration successfully")
}
