package category

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ninosistemas10/kiosko/infrastructure/postgres"
	"github.com/ninosistemas10/kiosko/model"
)

const table = "category"

var fields = []string{
	"id",
	"nombre",
	"description",
	"images",
	"activo",
	"created_at",
	"updated_at",
}

var (
	psqlInsert      = postgres.BuildSQLInsert(table, fields)
	psqlUpdate      = postgres.BuildSQLUpdateByID(table, fields)
	psqlDelete      = postgres.BuildSQLDelete(table)
	psqlGetAll      = postgres.BuildSQLSelect(table, fields)
	psqlUpdateImage = `UPDATE category SET images = $1, updated_at = $2 WHERE id = $3` // Nueva consulta
)

type Category struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Category {
	return Category{db}
}

func (c Category) Create(m *model.Category) error {
	_, err := c.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Nombre,
		m.Description,
		m.Images,
		m.Activo,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}
	return nil
}

func (c Category) Update(m *model.Category) error {
	_, err := c.db.Exec(
		context.Background(),
		psqlUpdate,
		m.Nombre,
		m.Description,
		m.Images,
		m.Activo,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c Category) UpdateImage(ID uuid.UUID, imagePath string) error {

	// Ejecutar la consulta de actualización
	_, err := c.db.Exec(
		context.Background(),
		psqlUpdateImage,
		imagePath,
		time.Now().Unix(),
		ID,
	)
	if err != nil {
		fmt.Println("❌ Error al actualizar la imagen en la base de datos:", err)
		return err
	}

	fmt.Println("✅ Imagen actualizada correctamente en la base de datos")
	return nil
}

// func (c Category) UpdateEsceptImage(ID uuid.UUID, updatedCategory model.Category) error {
// 	// Implementa la lógica para actualizar todos los campos excepto la imagen
// 	// Actualizar todos los campos excepto la imagen
// 	_, err := c.db.Exec(
// 		context.Background(),
// 		psqlUpdate,
// 		updatedCategory.Nombre,
// 		updatedCategory.Description,
// 		updatedCategory.Activo,
// 		updatedCategory.UpdatedAt,
// 		ID,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (c Category) Delete(ID uuid.UUID) error {
	_, err := c.db.Exec(
		context.Background(),
		psqlDelete,
		ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c Category) GetByID(ID uuid.UUID) (model.Category, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := c.db.QueryRow(
		context.Background(),
		query,
		ID,
	)
	return c.scanRow(row)
}

func (c Category) GetAll() (model.Categorys, error) {
	rows, err := c.db.Query(
		context.Background(),
		psqlGetAll,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.Categorys
	for rows.Next() {
		m, err := c.scanRow(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil

}

func (c Category) scanRow(s pgx.Row) (model.Category, error) {
	m := model.Category{}
	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Nombre,
		&m.Description,
		&m.Images,
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
