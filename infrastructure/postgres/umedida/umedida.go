package umedida

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ninosistemas10/kiosko/infrastructure/postgres"
	"github.com/ninosistemas10/kiosko/model"
)

const table = "unidadMedida"

var fields = []string{
	"id",
	"nombre",
	"abreviatura",
	"activo",
	"create_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, fields)
	psqlDelete = postgres.BuildSQLDelete(table)
)

type UnidadMedida struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) UnidadMedida {
	return UnidadMedida{db}
}

func (u UnidadMedida) Create(m *model.UnidadMedida) error {
	_, err := u.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Nombre,
		m.Abreviatura,
		m.Activo,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u UnidadMedida) Update(m *model.UnidadMedida) error {
	_, err := u.db.Exec(
		context.Background(),
		psqlUpdate,
		m.Nombre,
		m.Abreviatura,
		m.Activo,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u UnidadMedida) Delete(ID uuid.UUID) error {
	_, err := u.db.Exec(
		context.Background(),
		psqlDelete,
		ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u UnidadMedida) scanRow(s pgx.Row) (model.UnidadMedida, error) {
	m := model.UnidadMedida{}
	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Nombre,
		&m.Abreviatura,
		&m.Activo,
		&m.CreatedAt,
		&updatedAtNull,
	)

	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Int64

	return m, nil

}
