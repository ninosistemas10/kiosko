package handler

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/ninosistemas10/kiosko/infrastructure/handler/category"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/login"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/producto"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/subproducto"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/umedida"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/user"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/ws"
)

func InitRoutes(e *echo.Echo, dbPool *pgxpool.Pool) {
	// Ruta raíz
	e.Match([]string{"GET", "HEAD"}, "/", func(c echo.Context) error { // ✅ Permite ambos métodos
		return c.String(http.StatusOK, "¡Servidor en funcionamiento! ✅")
	})

	// Health check
	health(e)

	// WebSocket
	hub := ws.NewHub()
	go hub.Run()
	e.GET("/ws", func(c echo.Context) error { return ws.ServeWS(hub, c) })

	// Resto de rutas
	category.NewRouter(e, dbPool)
	login.NewRouter(e, dbPool)

	producto.NewRouter(e, dbPool)

	user.NewRouter(e, dbPool)
	umedida.NewRouter(e, dbPool)
	subproducto.NewRouter(e, dbPool)
}

func health(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			map[string]string{
				"time":         time.Now().String(),
				"message":      "Hello World!",
				"service_name": "",
			},
		)
	})
}
