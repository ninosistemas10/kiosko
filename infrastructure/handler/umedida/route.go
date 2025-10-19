package umedida

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/kiosko/domain/umedida"

	"github.com/ninosistemas10/kiosko/infrastructure/handler/middle"
	umedidaStorage "github.com/ninosistemas10/kiosko/infrastructure/postgres/umedida"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()

	adminRoutes(e, h, authMiddleware.IsValid, authMiddleware.IsAdmin)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := umedida.New(umedidaStorage.New(dbPool))
	return newHandler(useCase)
}

func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/ninosistemas/admin/umedida", middlewares...)

	route.POST("/", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)
}

func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/ninosistemas/public/umedida")

	route.POST("/", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)
}
