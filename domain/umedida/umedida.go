package umedida

import (
	"github.com/google/uuid"
	"github.com/ninosistemas10/kiosko/model"
)

type UseCase interface {
	Create(m *model.UnidadMedida) error
	Update(m *model.UnidadMedida) error
	Delete(ID uuid.UUID) error
}

type Storage interface {
	Create(m *model.UnidadMedida) error
	Update(m *model.UnidadMedida) error
	Delete(ID uuid.UUID) error
}
