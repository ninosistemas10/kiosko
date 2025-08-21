package model

import (
	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID `json:"id"`
	Nombre      string    `json:"nombre"`
	Description string    `json:"description"`
	Images      string    `json:"images"`
	//Images		json.RawMessage	`json:"images"`
	Activo    bool  `json:"activo"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func (c Category) HasID() bool { return c.ID != uuid.Nil }

type Categorys []Category

func (c Categorys) IsEmpty() bool { return len(c) == 0 }
