package main

import (
	"github.com/Amannigam1820/hr-dashboard-golang/config"
	"github.com/Amannigam1820/hr-dashboard-golang/database"
	"github.com/Amannigam1820/hr-dashboard-golang/router"
	"github.com/gofiber/fiber/v2"
)

func init() {
	database.ConnectDB()
}

func main() {
	config.InitCloudinary()

	sqlDb, err := database.DBConn.DB()
	if err != nil {
		panic("Error in sql connection")
	}
	defer sqlDb.Close()

	app := fiber.New()

	router.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": " qwerty Welcome to my first api in fiber"})

	})

	app.Listen(":8080")
}
