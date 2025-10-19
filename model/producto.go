package model

import (
	"github.com/google/uuid"
)

type Producto struct {
	ID              uuid.UUID `json:"id"`
	IdCategoria     uuid.UUID `json:"idCategoria"`
	IdUnidadMedida  uuid.UUID `json:"idUnidadMedida"`
	Nombre          string    `json:"nombre"`
	PrecioVenta     float64   `json:"precioVenta"`
	CostoPromedio   float64   `json:"costoPromedio"`
	StockMnimo      int32     `json:"stockMinimo"`
	Descripcion     string    `json:"descripcion"`
	Imagen          string    `json:"imagen"`
	Activo          bool      `json:"activo"`
	Destacado       bool      `json:"destacado"`
	LlevaInventario bool      `json:"llevaInventario"`
	CreateAt        int64     `json:"created_at"`
	UpdateAt        int64     `json:"updated_at"`
}

func (p Producto) HasID() bool {
	return p.ID != uuid.Nil
}

type Productos []Producto

func (p Productos) IsEmpty() bool {
	return len(p) == 0
}
