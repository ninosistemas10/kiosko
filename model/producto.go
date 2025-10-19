package model

import (
	"github.com/google/uuid"
)

type Producto struct {
	ID              uuid.UUID `json:"id"`
	IdCategoria     uuid.UUID `json:"idcategoria"`
	IdUnidadMedida  uuid.UUID `json:"idunidadmedida"`
	Nombre          string    `json:"nombre"`
	PrecioVenta     float64   `json:"precioventa"`
	CostoPromedio   float64   `json:"costopromedio"`
	StockMnimo      int32     `json:"stockminimo"`
	Descripcion     string    `json:"descripcion"`
	Imagen          string    `json:"imagen"`
	Activo          bool      `json:"activo"`
	Destacado       bool      `json:"destacado"`
	LlevaInventario bool      `json:"llevainventario"`
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
