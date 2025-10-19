package subproducto

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ninosistemas10/kiosko/model"
)

type SubProducto struct {
	storage Storage
}

func New(s Storage) SubProducto {
	return SubProducto{storage: s}
}

func (c SubProducto) Create(m *model.SubProducto) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}
	m.ID = ID

	if len(m.Imagen) == 0 {
		m.Imagen = ""
	}

	m.CreatedAt = time.Now().Unix()

	err = c.storage.Create(m)
	if err != nil {
		return err
	}
	return nil
}

func (c SubProducto) Update(m *model.SubProducto) error {
	if !m.HasID() {
		return fmt.Errorf("Update HasID")
	}

	//if len(m.Images) == 0 { m.Images = []byte(`{}`) }
	if len(m.Imagen) == 0 {
		m.Imagen = ""
	}

	m.UpdatedAt = time.Now().Unix()

	err := c.storage.Update(m)
	if err != nil {
		return err
	}

	return nil
}

func (c SubProducto) UpdateImage(ID uuid.UUID, imagePath string) error {
	// Verificar si el ID es válido
	if ID == uuid.Nil {
		return fmt.Errorf("invalid ID")
	}

	// Intentar actualizar la imagen en la base de datos
	err := c.storage.UpdateImage(ID, imagePath)
	if err != nil {
		return err
	}

	return nil
}

func (c SubProducto) Delete(ID uuid.UUID) error {
	err := c.storage.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}

func (c SubProducto) GetByID(ID uuid.UUID) (model.SubProducto, error) {
	subProducto, err := c.storage.GetByID(ID)
	if err != nil {
		return model.SubProducto{}, fmt.Errorf("SubProducto: %w", err)
	}

	return subProducto, nil
}

func (c SubProducto) GetAll() (model.SubProductos, error) {
	subProductos, err := c.storage.GetAll()
	if err != nil {
		return model.SubProductos{}, fmt.Errorf("SubProductos: %w", err)
	}

	return subProductos, nil
}
