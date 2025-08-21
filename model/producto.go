package model

import (
	"time"

	"github.com/google/uuid"
)

type Producto struct {
	ID          uuid.UUID `json:"id"`
	IdCategoria uuid.UUID `json:"idcategoria"`
	Nombre      string    `json:"nombre"`
	Precio      float64   `json:"precio"`
	Imagen      string    `json:"imagen"`
	Descripcion string    `json:"descripcion"`
	Activo      bool      `json:"activo"`
	Time        time.Time `json:"time"`
	Calorias    float64   `json:"calorias"`
	CreateAt    int64     `json:"created_at"`
	UpdateAt    int64     `json:"updated_at"`
}

func (p Producto) HasID() bool {
	return p.ID != uuid.Nil
}

type Productos []Producto

func (p Productos) IsEmpty() bool {
	return len(p) == 0
}
