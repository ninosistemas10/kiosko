package category

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/kiosko/domain/category"

	"github.com/ninosistemas10/kiosko/infrastructure/handler/middle"
	categoryStorage "github.com/ninosistemas10/kiosko/infrastructure/postgres/category"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := builHandler(dbPool)

	authMiddlleware := middle.New()

	adminRoutes(e, h, authMiddlleware.IsValid, authMiddlleware.IsAdmin)
	publicRoutes(e, h)
}

func builHandler(dbPool *pgxpool.Pool) handler {
	useCase := category.New(categoryStorage.New(dbPool))
	return newHandler(useCase)
}

func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("ninosistemas/admin/category", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
}

func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("ninosistemas/public/category")

	route.POST("", h.Create)
	route.PUT("/imagen/:id", h.UpdateImage)
	route.GET("", h.GetAll)
	route.DELETE("/:id", h.Delete)
	route.PUT("/:id", h.Update)
}
