package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

type PostgresMaterialRepository struct {
	db *pgxpool.Pool
}

func NewPostgresMaterialRepository(db *pgxpool.Pool) *PostgresMaterialRepository {
	return &PostgresMaterialRepository{db: db}
}

func (r *PostgresMaterialRepository) FindAll() ([]entity.DimMaterial, error) {
	query := `
		SELECT
			id_material::text,
			id_material::text,
			codigo_material,
			descricao,
			COALESCE(categoria, ''),
			COALESCE(fabricante, ''),
			COALESCE(custo_estimado::text, ''),
			COALESCE(status, '')
		FROM material
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DimMaterial
	for rows.Next() {
		var m entity.DimMaterial
		err := rows.Scan(
			&m.SkMaterial, &m.IdMaterial, &m.CodigoMaterial, &m.Descricao,
			&m.Categoria, &m.Fabricante, &m.CustoEstimado, &m.Status,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, m)
	}

	return results, rows.Err()
}
