package subproducto

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

const table = "subproductos"

var fields = []string{
	"id",
	"idproducto",
	"nombre",
	"precioadicional",
	"imagen",
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

type SubProducto struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) SubProducto {
	return SubProducto{db}
}

func (c SubProducto) Create(m *model.SubProducto) error {
	_, err := c.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.IdProducto,
		m.Nombre,
		m.PrecioAdicional,
		m.Imagen,
		m.Activo,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}
	return nil
}

func (c SubProducto) Update(m *model.SubProducto) error {
	_, err := c.db.Exec(
		context.Background(),
		psqlUpdate,
		m.ID,
		m.IdProducto,
		m.Nombre,
		m.PrecioAdicional,
		m.Imagen,
		m.Activo,
		m.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c SubProducto) UpdateImage(ID uuid.UUID, imagePath string) error {

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

func (c SubProducto) Delete(ID uuid.UUID) error {
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

func (c SubProducto) GetByID(ID uuid.UUID) (model.SubProducto, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := c.db.QueryRow(
		context.Background(),
		query,
		ID,
	)
	return c.scanRow(row)
}

func (c SubProducto) GetAll() (model.SubProductos, error) {
	rows, err := c.db.Query(
		context.Background(),
		psqlGetAll,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.SubProductos
	for rows.Next() {
		m, err := c.scanRow(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil

}

func (c SubProducto) scanRow(s pgx.Row) (model.SubProducto, error) {
	m := model.SubProducto{}
	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.IdProducto,
		&m.Nombre,
		&m.PrecioAdicional,
		&m.Imagen,
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
