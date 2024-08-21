package main

import (
	"e-wallet/database"
	"e-wallet/database/migration"
	"e-wallet/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//Database Initial
	database.DatabaseInit()

	// Database Migration
	migration.Migration()

	// Inisialisasi Route
	routes.RouteInit(app)

	app.Listen(":3000")
}
