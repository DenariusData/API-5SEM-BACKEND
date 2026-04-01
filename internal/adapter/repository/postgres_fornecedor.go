package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

type PostgresFornecedorRepository struct {
	db *pgxpool.Pool
}

func NewPostgresFornecedorRepository(db *pgxpool.Pool) *PostgresFornecedorRepository {
	return &PostgresFornecedorRepository{db: db}
}

func (r *PostgresFornecedorRepository) FindAll() ([]entity.DimFornecedor, error) {
	query := `
		SELECT
			id_fornecedor::text,
			id_fornecedor::text,
			codigo_fornecedor,
			razao_social,
			COALESCE(cidade, ''),
			COALESCE(estado, ''),
			COALESCE(categoria, ''),
			COALESCE(status, '')
		FROM fornecedor
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DimFornecedor
	for rows.Next() {
		var f entity.DimFornecedor
		err := rows.Scan(
			&f.SkFornecedor, &f.IdFornecedor, &f.CodigoFornecedor, &f.RazaoSocial,
			&f.Cidade, &f.Estado, &f.Categoria, &f.Status,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, f)
	}

	return results, rows.Err()
}
