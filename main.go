// main.go
package main

import (
	"log"
	"postgress_api/database"
	"postgress_api/route"
	"github.com/gofiber/fiber/v2"
)


func main() {

	// Create a new Fiber instance
	app := fiber.New()

	// Initialize Database
	database.ConnectDB()

	// Setup routes
	route.UserRoutes(app)

	// Start the Fiber app
	log.Fatal(app.Listen(":3000"))
	
}
