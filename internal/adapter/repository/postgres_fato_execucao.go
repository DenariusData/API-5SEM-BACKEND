package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

type PostgresFatoExecucaoRepository struct {
	db *pgxpool.Pool
}

func NewPostgresFatoExecucaoRepository(db *pgxpool.Pool) *PostgresFatoExecucaoRepository {
	return &PostgresFatoExecucaoRepository{db: db}
}

func (r *PostgresFatoExecucaoRepository) FindAll() ([]entity.FatoExecucaoTarefas, error) {
	query := `
		SELECT
			tt.id_tempo::text,
			t.id_projeto::text,
			tt.id_tarefa::text,
			COALESCE(tt.usuario, ''),
			tt.data::text,
			tt.horas_trabalhadas::text
		FROM tempo_tarefa tt
		JOIN tarefa t ON tt.id_tarefa = t.id_tarefa
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.FatoExecucaoTarefas
	for rows.Next() {
		var f entity.FatoExecucaoTarefas
		err := rows.Scan(
			&f.SkFato, &f.SkProjeto, &f.SkTarefa, &f.SkResponsavel,
			&f.SkTempo, &f.HorasTrabalhadas,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, f)
	}

	return results, rows.Err()
}
