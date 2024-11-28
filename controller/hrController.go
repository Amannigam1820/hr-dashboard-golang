package controller

import (
	"github.com/Amannigam1820/hr-dashboard-golang/database"
	"github.com/Amannigam1820/hr-dashboard-golang/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateHr(c *fiber.Ctx) error {
	var hr model.Hr
	if err := c.BodyParser(&hr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": "false",
			"error":   "Failed to parse request body",
		})
	}

	if hr.Name == "" || hr.Email == "" || hr.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": "false",
			"error":   "Name and Email and Password are required fields",
		})
	}

	var existingHr model.Hr
	if err := database.DBConn.Where("email = ?", hr.Email).First(&existingHr).Error; err == nil {
		// If no error, it means the email already exists
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": "false",
			"message": "Email already exists",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(hr.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": "false",
			"error":   "Failed to hash password",
		})
	}
	hr.Password = string(hashedPassword)

	result := database.DBConn.Create(&hr)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": "false",
			"error":   "Failed to create HR record",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "true",
		"message": "Hr Created SuccessFully",
		"Hr":      hr,
	})
}

func GetAllHr(c *fiber.Ctx) error {
	var hrs []model.Hr
	if result := database.DBConn.Find(&hrs); result.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to retrieve HR records",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"data":    hrs,
		"success": true,
	})
}
