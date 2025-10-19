package subproducto

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/middle"

	"github.com/ninosistemas10/kiosko/domain/subproducto"
	subProductoStorage "github.com/ninosistemas10/kiosko/infrastructure/postgres/subproducto"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()

	adminRoutes(e, h, authMiddleware.IsValid, authMiddleware.IsAdmin)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := subproducto.New(subProductoStorage.New(dbPool))
	return newHandler(useCase)
}

func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/ninosistemas/admin/subProductos", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAll)
	//route.GET("/:id", h.GetByID)
}

func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/ninosistemas/public/subProductos")

	route.POST("", h.Create)
	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
	route.PUT("/imagen/:id", h.UpdateImage)
	//route.GET("/categoria/:idcategoria", h.GetBySubProductosID)
}
