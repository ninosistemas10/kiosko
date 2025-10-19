package model

import "github.com/google/uuid"

type SubProducto struct {
	ID              uuid.UUID `json:"id"`
	IdProducto      uuid.UUID `json:"idproducto"`
	Nombre          string    `json:"nombre"`
	PrecioAdicional float64   `json:"precioadicional"`
	Imagen          string    `json:"imagen"`
	Activo          bool      `json:"activo"`
	CreatedAt       int64     `json:"created_at"`
	UpdatedAt       int64     `json:"updated_at"`
}

func (s SubProducto) HasID() bool {
	return s.ID != uuid.Nil
}

type SubProductos []SubProducto

func (s SubProductos) IsEmpty() bool { return len(s) == 0 }
