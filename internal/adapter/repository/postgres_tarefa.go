package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

type PostgresTarefaRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTarefaRepository(db *pgxpool.Pool) *PostgresTarefaRepository {
	return &PostgresTarefaRepository{db: db}
}

func (r *PostgresTarefaRepository) FindAll() ([]entity.DimTarefa, error) {
	query := `
		SELECT
			id_tarefa::text,
			id_tarefa::text,
			codigo_tarefa,
			titulo,
			COALESCE(status, ''),
			COALESCE(estimativa_horas::text, ''),
			COALESCE(data_inicio::text, ''),
			COALESCE(data_fim_prevista::text, '')
		FROM tarefa
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DimTarefa
	for rows.Next() {
		var t entity.DimTarefa
		err := rows.Scan(
			&t.SkTarefa, &t.IdTarefa, &t.CodigoTarefa, &t.Titulo,
			&t.Status, &t.EstimativaHoras, &t.DataInicio, &t.DataFimPrevista,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, t)
	}

	return results, rows.Err()
}
