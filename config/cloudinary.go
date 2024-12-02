package config

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var CloudinaryClient *cloudinary.Cloudinary

func InitCloudinary() {
	cld, err := cloudinary.NewFromParams("dcohtcgb7", "997959291724572", "RYS6MeK3gp5qCG3OfSj99j_wwuw")
	if err != nil {
		log.Fatal("Error initializing Cloudinary: ", err)
	}
	CloudinaryClient = cld
}

func UploadToCloudinary(file multipart.File) (string, error) {
	if CloudinaryClient == nil {
		return "", fmt.Errorf("Cloudinary client is not initialized")
	}
	// Upload to Cloudinary using the latest API
	uploadResult, err := CloudinaryClient.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder: "employee_data", // Optional folder name in Cloudinary
	})
	if err != nil {
		return "", fmt.Errorf("error uploading to Cloudinary: %v", err)
	}
	return uploadResult.SecureURL, nil
}
