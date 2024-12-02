package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/Amannigam1820/hr-dashboard-golang/config"
	"github.com/Amannigam1820/hr-dashboard-golang/database"
	"github.com/Amannigam1820/hr-dashboard-golang/model"
	"github.com/gofiber/fiber/v2"
)

func CreateEmployee(c *fiber.Ctx) error {
	var employee model.Employee

	// Parse multipart form data (including file upload)
	err := c.BodyParser(&employee)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unable to parse request body")
	}

	// Handle file uploads
	form, err := c.MultipartForm()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Unable to parse multipart form")
	}

	// Process files (Resume, Experience Letter, Relieving Letter)
	if len(form.File["resume"]) > 0 {
		file := form.File["resume"][0] // *multipart.FileHeader
		f, err := file.Open()          // Open the file to get a multipart.File (the content of the file)
		if err != nil {
			log.Println("Error opening resume file:", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to open resume file")
		}
		url, err := config.UploadToCloudinary(f) // Pass the file content to UploadToCloudinary
		if err != nil {
			log.Println("Error uploading resume:", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload resume")
		}
		employee.Resume = url
		defer f.Close() // Don't forget to close the file after usage
	}

	if len(form.File["experience_letter"]) > 0 {
		file := form.File["experience_letter"][0] // *multipart.FileHeader
		f, err := file.Open()                     // Open the file to get a multipart.File (the content of the file)
		if err != nil {
			log.Println("Error opening experience letter file:", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to open experience letter file")
		}
		url, err := config.UploadToCloudinary(f) // Pass the file content to UploadToCloudinary
		if err != nil {
			log.Println("Error uploading experience letter:", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload experience letter")
		}
		employee.ExperienceLetter = url
		defer f.Close() // Don't forget to close the file after usage
	}

	if len(form.File["releiving_letter"]) > 0 {
		file := form.File["releiving_letter"][0] // *multipart.FileHeader
		f, err := file.Open()                    // Open the file to get a multipart.File (the content of the file)
		if err != nil {
			log.Println("Error opening relieving letter file:", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to open relieving letter file")
		}
		url, err := config.UploadToCloudinary(f) // Pass the file content to UploadToCloudinary
		if err != nil {
			log.Println("Error uploading relieving letter:", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload relieving letter")
		}
		employee.ReleivingLetter = url
		defer f.Close() // Don't forget to close the file after usage
	}

	// Set the created date
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()

	fmt.Println(&employee.CreatedAt)
	fmt.Println(&employee.BirthDate)

	// Insert employee into the database
	result := database.DBConn.Create(&employee)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": "false",
			"error":   "Failed to create HR record",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "true",
		"message": "Hr Created SuccessFully",
		"Hr":      employee,
	})

}

func GetAllEmployee(c *fiber.Ctx) error {
	var employees []model.Employee

	if result := database.DBConn.Find(&employees); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to retrieve Employee records",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"data":    employees,
		"success": true,
	})
}

func GetEmployeeById(c *fiber.Ctx) error {
	var employee model.Employee
	id := c.Params("id")

	if err := database.DBConn.First(&employee, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "HR record not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"hr":      employee,
	})
}
