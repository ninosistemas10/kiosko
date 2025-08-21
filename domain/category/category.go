package category

import (
	"github.com/google/uuid"
	"github.com/ninosistemas10/kiosko/model"
)

type UseCase interface {
	Create(m *model.Category) error
	Update(m *model.Category) error
	UpdateImage(ID uuid.UUID, imagePath string) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Category, error)
	GetAll() (model.Categorys, error)
}

type Storage interface {
	Create(m *model.Category) error
	Update(m *model.Category) error
	UpdateImage(ID uuid.UUID, imagePath string) error
	Delete(ID uuid.UUID) error

	GetByID(ID uuid.UUID) (model.Category, error)
	GetAll() (model.Categorys, error)
}
