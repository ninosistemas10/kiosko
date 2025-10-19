package subproducto

import (
	"github.com/google/uuid"
	"github.com/ninosistemas10/kiosko/model"
)

type UseCase interface {
	Create(m *model.SubProducto) error
	Update(m *model.SubProducto) error
	UpdateImage(ID uuid.UUID, imagePath string) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.SubProducto, error)
	GetAll() (model.SubProductos, error)
}

type Storage interface {
	Create(m *model.SubProducto) error
	Update(m *model.SubProducto) error
	UpdateImage(ID uuid.UUID, imagePath string) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.SubProducto, error)
	GetAll() (model.SubProductos, error)
}
