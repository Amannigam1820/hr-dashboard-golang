package router

import (
	"github.com/Amannigam1820/hr-dashboard-golang/controller"
	"github.com/Amannigam1820/hr-dashboard-golang/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupHRRoutes sets up the routes for HR related operations
func SetupHRRoutes(app *fiber.App) {
	hrGroup := app.Group("/api/hr")

	// Define the HR routes
	hrGroup.Post("/", middleware.RoleCheck("Super-Admin"), controller.CreateHr)
	hrGroup.Get("/all", controller.GetAllHr)
	hrGroup.Get("/:id", middleware.RoleCheck("Super-Admin"), controller.GetHrById)
	hrGroup.Put("/:id", middleware.RoleCheck("Super-Admin"), controller.UpdateHr)
	hrGroup.Delete("/:id", middleware.RoleCheck("Super-Admin"), controller.DeleteHr)
	hrGroup.Post("/login", controller.LoginHr)
	hrGroup.Post("/logout", controller.Logout)
}
