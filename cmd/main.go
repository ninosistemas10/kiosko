package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/kiosko/infrastructure/handler"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/response"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/ws"
)

func main() {
	// Cargara variables de entorno
	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}
	//validar variables de entorno
	err := validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	//crea servidor HTTP (ECHO O TU wrapper)
	e := newHTTP(response.HTTPErrorHandler)
	// Crear instancia de dbPool
	//Crea la base de Datos
	dbPool, err := newDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	//Inicializa las rutas HTTP
	handler.InitRoutes(e, dbPool)

	hub := ws.NewHub()
	go hub.Run()   //Inicia el hub en un goroytine
	ws.SetHub(hub) // Guardar hub globakl para Emit()
	e.GET("/ws", func(c echo.Context) error { return ws.ServeWS(hub, c) })

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081" // fallback local
	}
	err = e.Start(":" + port)
	if err != nil {
		log.Fatal(err)
	}

}
