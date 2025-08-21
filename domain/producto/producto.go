package producto

import (
	"github.com/google/uuid"
	"github.com/ninosistemas10/kiosko/model"
)

type UseCase interface {
	Create(m *model.Producto) error
	Update(m *model.Producto) error
	Delete(ID uuid.UUID) error

	GetByCategoryID(idCategoria uuid.UUID) (model.Productos, error)
	GetByID(ID uuid.UUID) (model.Producto, error)
	GetAll() (model.Productos, error)

	UpdateImage(ID uuid.UUID, imagePath string) error
}

type Storage interface {
	Create(m *model.Producto) error
	Update(m *model.Producto) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Producto, error)
	GetByCategoryID(idCategoria uuid.UUID) (model.Productos, error)
	GetAll() (model.Productos, error)
	UpdateImage(ID uuid.UUID, imagePath string) error
}
