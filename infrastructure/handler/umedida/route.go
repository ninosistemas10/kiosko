package umedida

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/kiosko/domain/umedida"

	"github.com/ninosistemas10/kiosko/infrastructure/handler/middle"
	umedidaStorage "github.com/ninosistemas10/kiosko/infrastructure/postgres/umedida"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := builHandler(dbPool)

	authMiddlleware := middle.New()

	adminRoutes(e, h, authMiddlleware.IsValid, authMiddlleware.IsAdmin)
	publicRoutes(e, h)
}

func builHandler(dbPool *pgxpool.Pool) handler {
	useCase := umedida.New(umedidaStorage.New(dbPool))
	return newHandler(useCase)
}

func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/ninosistemas/admin/umedida", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

}

func publicRoutes(e *echo.Echo, h handler) {

	e.POST("/ninosistemas/public/umedida", h.Create)
	e.PUT("/ninosistemas/public/umedida/:id", h.Update)
	e.DELETE("/ninosistemas/public/umedida/:id", h.Delete)
}
