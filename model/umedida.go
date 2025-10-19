package model

import "github.com/google/uuid"

type UnidadMedida struct {
	ID          uuid.UUID `json:"id"`
	Nombre      string    `json:"nombre"`
	Abreviatura string    `json:"abreviatura"`
	Activo      bool      `json:"activo"`
	CreatedAt   int64     `json:"created_at"`
	UpdatedAt   int64     `json:"updated_at"`
}

func (u UnidadMedida) HasID() bool { return u.ID != uuid.Nil }
