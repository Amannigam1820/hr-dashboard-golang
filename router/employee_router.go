package router

import (
	"github.com/Amannigam1820/hr-dashboard-golang/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupEmployeeRoutes(app *fiber.App) {
	employeeGroup := app.Group("/api/employee")

	employeeGroup.Post("/", controller.CreateEmployee)
	employeeGroup.Get("/all", controller.GetAllEmployee)
	employeeGroup.Get("/:id", controller.GetEmployeeById)
	employeeGroup.Delete("/:id", controller.DeleteEmployee)
	employeeGroup.Put("/:id", controller.UpdateEmployee)

}
