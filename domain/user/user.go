package user

import (
	"github.com/google/uuid"
	"github.com/ninosistemas10/kiosko/model"
)

type UseCase interface {
	Create(m *model.User) error
	GetByID(ID uuid.UUID) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

type Storage interface {
	Create(m *model.User) error
	GetByID(ID uuid.UUID) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}
