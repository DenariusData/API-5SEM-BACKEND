package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

type PostgresResponsavelRepository struct {
	db *pgxpool.Pool
}

func NewPostgresResponsavelRepository(db *pgxpool.Pool) *PostgresResponsavelRepository {
	return &PostgresResponsavelRepository{db: db}
}

func (r *PostgresResponsavelRepository) FindAll() ([]entity.DimResponsavel, error) {
	query := `
		SELECT
			ROW_NUMBER() OVER (ORDER BY nome_responsavel)::text,
			nome_responsavel,
			tipo
		FROM (
			SELECT DISTINCT responsavel AS nome_responsavel, 'Projeto' AS tipo FROM projeto WHERE responsavel IS NOT NULL
			UNION
			SELECT DISTINCT responsavel, 'Tarefa' FROM tarefa WHERE responsavel IS NOT NULL
		) r
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DimResponsavel
	for rows.Next() {
		var resp entity.DimResponsavel
		err := rows.Scan(&resp.SkResponsavel, &resp.NomeResponsavel, &resp.Tipo)
		if err != nil {
			return nil, err
		}
		results = append(results, resp)
	}

	return results, rows.Err()
}
