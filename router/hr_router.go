package router

import (
	"github.com/Amannigam1820/hr-dashboard-golang/controller"
	"github.com/gofiber/fiber/v2"
)

// SetupHRRoutes sets up the routes for HR related operations
func SetupHRRoutes(app *fiber.App) {
	hrGroup := app.Group("/api/hr")

	// Define the HR routes
	hrGroup.Post("/", controller.CreateHr)
	hrGroup.Get("/all", controller.GetAllHr)
}
