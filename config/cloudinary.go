package config

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go"
)

func SetupCloudinary() *cloudinary.Cloudinary {
	clr, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		log.Fatalf("Cloudinary configure failed %v", err)
	}
	return clr
}
