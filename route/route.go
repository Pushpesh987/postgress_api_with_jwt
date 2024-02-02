// route/route.go
package route

import (
	"postgress_api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Routes
func UserRoutes(app *fiber.App) {
	
	
	// Middleware
	api := app.Group("/api", logger.New())

	// login
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// signup
	user := auth.Group("/signup")
	user.Post("/", handler.SignUp)
}