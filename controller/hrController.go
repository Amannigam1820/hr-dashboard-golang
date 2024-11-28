package controller

import (
	"time"

	"github.com/Amannigam1820/hr-dashboard-golang/database"
	"github.com/Amannigam1820/hr-dashboard-golang/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

func GetHrById(c *fiber.Ctx) error {
	var hr model.Hr
	id := c.Params("id")
	if err := database.DBConn.First(&hr, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "HR record not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"hr":      hr,
	})

}

func UpdateHr(c *fiber.Ctx) error {
	var hr model.Hr
	id := c.Params("id")

	if err := c.BodyParser(&hr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to parse request body",
		})
	}
	var existingHr model.Hr

	if result := database.DBConn.First(&existingHr, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Hr Record not found",
		})
	}
	if hr.Name != "" {
		existingHr.Name = hr.Name
	}

	if hr.Email != "" {
		existingHr.Email = hr.Email

		var existingEmail model.Hr
		if err := database.DBConn.Where("email:?", hr.Email).First(&existingEmail).Error; err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Email already exists",
			})
		}
	}
	if hr.Password != "" {
		// If the password is provided, hash it before saving
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(hr.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to hash password",
			})
		}
		existingHr.Password = string(hashedPassword)
	}

	if result := database.DBConn.Save(&existingHr); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to update HR record",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "HR record updated successfully !",
		"data":    existingHr,
	})

}

func DeleteHr(c *fiber.Ctx) error {
	id := c.Params("id")
	var hr model.Hr

	if err := database.DBConn.First(&hr, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "HR record not found",
		})
	}
	if result := database.DBConn.Delete(&hr); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to delete HR record",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "HR record deleted successfully",
	})
}

func LoginHr(c *fiber.Ctx) error {
	var hr model.Hr
	if err := c.BodyParser(&hr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to parse request body",
		})
	}
	var existingHr model.Hr
	if err := database.DBConn.Where("email =  ?", hr.Email).First(&existingHr).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid email or password",
		})
	}
	err := bcrypt.CompareHashAndPassword([]byte(existingHr.Password), []byte(hr.Password))
	if err != nil {
		// If password doesn't match, return an error
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid email or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    existingHr.ID,
		"email": existingHr.Email,
		"exp":   time.Now().Add(time.Hour * 24),
	})

	secretKey := []byte("hrdashboard")

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to generate token",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"token":   tokenString,
	})

}
func Logout(c *fiber.Ctx) error {
	// Clear the JWT cookie by setting its value to an empty string and setting its expiration time to the past
	c.Cookie(&fiber.Cookie{
		Name:     "token",                        // Name of the cookie
		Value:    "",                             // Set the value to empty
		Expires:  time.Now().Add(-1 * time.Hour), // Expire the cookie in the past to remove it
		HTTPOnly: true,                           // Keep the cookie HTTPOnly for security
		Secure:   false,                          // Set to true if using HTTPS
		SameSite: "Strict",                       // Prevent the cookie from being sent in cross-site requests
	})

	// Return a success response indicating the user has logged out
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Logged out successfully",
	})
}
