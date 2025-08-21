package main

import (
	"log"
	"os"

	"github.com/ninosistemas10/kiosko/infrastructure/handler"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/response"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}

	err := validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	e := newHTTP(response.HTTPErrorHandler)

	dbPool, err := newDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoutes(e, dbPool)

	err = e.Start(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}

}
