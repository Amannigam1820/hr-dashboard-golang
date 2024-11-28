package router

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all routes for the application
func SetupRoutes(app *fiber.App) {
	// Setup routes for HR and Employee
	SetupHRRoutes(app)
	//SetupHRRoutes(app)
	//SetupEmployeeRoutes(app)
}
