package umedida

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/kiosko/domain/umedida"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/response"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/ws"
	"github.com/ninosistemas10/kiosko/model"
)

type handler struct {
	useCase  umedida.UseCase
	response response.API
}

func newHandler(useCase umedida.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create(c echo.Context) error {
	m := model.UnidadMedida{}
	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	ws.Emit("unidadmedida", "created", m)
	return c.JSON(h.response.Created(m))
}

func (h handler) Update(c echo.Context) error {
	m := model.UnidadMedida{}
	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	m.ID = ID

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c, "h.useCase.Update()", err)
	}

	ws.Emit("unidadmedida", "updated", m)
	return c.JSON(h.response.Updated(m))
}

func (h handler) Delete(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	err = h.useCase.Delete(ID)
	if err != nil {
		return h.response.Error(c, "useCase.Delete()", err)
	}

	ws.Emit("umedida", "deleted", map[string]interface{}{"id": ID})
	return c.JSON(h.response.Deleted(nil))
}
