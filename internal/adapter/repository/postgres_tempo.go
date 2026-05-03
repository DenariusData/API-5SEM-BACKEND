package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

type PostgresTempoRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTempoRepository(db *pgxpool.Pool) *PostgresTempoRepository {
	return &PostgresTempoRepository{db: db}
}

func (r *PostgresTempoRepository) FindAll() ([]entity.DimTempo, error) {
	query := `
		SELECT
			id_tempo::text,
			data::text,
			EXTRACT(YEAR FROM data)::text,
			EXTRACT(MONTH FROM data)::text,
			EXTRACT(DAY FROM data)::text,
			TO_CHAR(data, 'Day'),
			EXTRACT(WEEK FROM data)::text,
			CASE WHEN EXTRACT(MONTH FROM data) <= 6 THEN '1' ELSE '2' END,
			EXTRACT(QUARTER FROM data)::text,
			TO_CHAR(data, 'TMMonth')
		FROM tempo_tarefa
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DimTempo
	for rows.Next() {
		var t entity.DimTempo
		err := rows.Scan(
			&t.SkTempo, &t.DataCompleta, &t.Ano, &t.Mes, &t.Dia,
			&t.DiaSemana, &t.Semana, &t.Semestre, &t.Trimestre, &t.NomeMes,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, t)
	}

	return results, rows.Err()
}

func (r *PostgresTempoRepository) GetTempoGasto() (interface{}, error) {
	query := `
		SELECT COALESCE(SUM(horas_trabalhadas), 0)
		FROM tempo_tarefa
	`

	var total float64

	err := r.db.QueryRow(context.Background(), query).Scan(&total)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"total_tempo_gasto": total,
	}

	return result, nil
}