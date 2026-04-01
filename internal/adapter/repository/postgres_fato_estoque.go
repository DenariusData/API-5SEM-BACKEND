package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

type PostgresFatoEstoqueRepository struct {
	db *pgxpool.Pool
}

func NewPostgresFatoEstoqueRepository(db *pgxpool.Pool) *PostgresFatoEstoqueRepository {
	return &PostgresFatoEstoqueRepository{db: db}
}

func (r *PostgresFatoEstoqueRepository) FindAll() ([]entity.FatoEstoqueMateriais, error) {
	query := `
		SELECT
			ep.id_estoque::text,
			ep.id_projeto::text,
			ep.id_material::text,
			COALESCE(e.data_empenho::text, ''),
			ep.quantidade::text,
			COALESCE(e.quantidade_empenhada::text, '0')
		FROM estoque_projeto ep
		LEFT JOIN (
			SELECT id_projeto, id_material,
			       MAX(data_empenho) AS data_empenho,
			       SUM(quantidade_empenhada) AS quantidade_empenhada
			FROM empenho
			GROUP BY id_projeto, id_material
		) e ON ep.id_projeto = e.id_projeto AND ep.id_material = e.id_material
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.FatoEstoqueMateriais
	for rows.Next() {
		var f entity.FatoEstoqueMateriais
		err := rows.Scan(
			&f.SkFato, &f.SkProjeto, &f.SkMaterial, &f.SkTempo,
			&f.QuantidadeEstoque, &f.QuantidadeEmpenhada,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, f)
	}

	return results, rows.Err()
}
