package main

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func validateEnvironments() error {
	if strings.TrimSpace(os.Getenv("SERVER_PORT")) == "" {
		return errors.New("the SERVER_PORT ")
	}
	if strings.TrimSpace(os.Getenv("ALLOWED_ORIGINS")) == "" {
		return errors.New("the ALLOWED_ORIGINS ")
	}
	if strings.TrimSpace(os.Getenv("ALLOWED_METHODS")) == "" {
		return errors.New("the ALLOWED_METHODS")
	}
	if strings.TrimSpace(os.Getenv("IMAGES_DIR")) == "" {
		return errors.New("the IMAGES_DIR")
	}

	if strings.TrimSpace(os.Getenv("DB_USER")) == "" {
		return errors.New("DB_USER is required")
	}
	if strings.TrimSpace(os.Getenv("DB_PASSWORD")) == "" {
		return errors.New("the ALLOWED_METHODS")
	}
	if strings.TrimSpace(os.Getenv("DB_HOST")) == "" {
		return errors.New("the ALLOWED_METHODS")
	}
	if strings.TrimSpace(os.Getenv("DB_PORT")) == "" {
		return errors.New("the ALLOWED_METHODS")
	}
	if strings.TrimSpace(os.Getenv("DB_NAME")) == "" {
		return errors.New("the ALLOWED_METHODS")
	}
	if strings.TrimSpace(os.Getenv("DB_SSL_MODE")) == "" {
		return errors.New("the env is mandatory")
	}

	// PayPal envs.
	if strings.TrimSpace(os.Getenv("WEBHOOK_ID")) == "" {
		return errors.New("the WEBHOOK_ID env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("VALIDATION_URL")) == "" {
		return errors.New("the VALIDATION_URL env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("CLIENT_ID")) == "" {
		return errors.New("the CLIENT_ID env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("SECRET_ID")) == "" {
		return errors.New("the SECRET_ID env is mandatory")
	}

	if strings.TrimSpace(os.Getenv("CLOUDINARY_CLOUD_NAME")) == "" {
		return errors.New("the WEBHOOK_ID env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("CLOUDINARY_API_KEY")) == "" {
		return errors.New("the VALIDATION_URL env is mandatory")
	}
	if strings.TrimSpace(os.Getenv("CLOUDINARY_API_SECRET")) == "" {
		return errors.New("the CLIENT_ID env is mandatory")
	}

	return nil
}
