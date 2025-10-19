package umedida

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ninosistemas10/kiosko/model"
)

type UnidadMedida struct {
	storage Storage
}

func New(s Storage) UnidadMedida {
	return UnidadMedida{storage: s}
}

func (u UnidadMedida) Create(m *model.UnidadMedida) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}
	m.ID = ID

	m.CreatedAt = time.Now().Unix()

	err = u.storage.Create(m)
	if err != nil {
		return err
	}
	return nil
}

func (u UnidadMedida) Update(m *model.UnidadMedida) error {
	if !m.HasID() {
		return fmt.Errorf("Update HasID")
	}

	m.UpdatedAt = time.Now().Unix()

	err := u.storage.Update(m)
	if err != nil {
		return err
	}

	return nil
}

func (u UnidadMedida) Delete(ID uuid.UUID) error {
	err := u.storage.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}
