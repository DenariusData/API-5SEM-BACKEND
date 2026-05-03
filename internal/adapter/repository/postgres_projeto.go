package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

type PostgresProjetoRepository struct {
	db *pgxpool.Pool
}

func NewPostgresProjetoRepository(db *pgxpool.Pool) *PostgresProjetoRepository {
	return &PostgresProjetoRepository{db: db}
}

func (r *PostgresProjetoRepository) FindAll() ([]entity.DimProjeto, error) {
	query := `
		SELECT
			p.id_projeto::text,
			p.id_projeto::text,
			p.codigo_projeto,
			p.nome_projeto,
			COALESCE(p.responsavel, ''),
			COALESCE(p.status, ''),
			COALESCE(pr.codigo_programa, ''),
			COALESCE(pr.nome_programa, ''),
			COALESCE(pr.gerente_programa, ''),
			COALESCE(p.data_inicio::text, ''),
			COALESCE(p.data_fim_prevista::text, '')
		FROM projeto p
		LEFT JOIN programa pr ON p.id_programa = pr.id_programa
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DimProjeto
	for rows.Next() {
		var p entity.DimProjeto
		err := rows.Scan(
			&p.SkProjeto, &p.IdProjeto, &p.CodigoProjeto, &p.NomeProjeto,
			&p.Responsavel, &p.Status, &p.CodigoPrograma, &p.NomePrograma,
			&p.GerentePrograma, &p.DataInicio, &p.DataFimPrevista,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, p)
	}

	return results, rows.Err()
}

func (r *PostgresProjetoRepository) FindInvestimentoByPrograma() ([]entity.ProgramaInvestimento, error) {
	query := `
		SELECT
			pr.codigo_programa,
			pr.nome_programa,
			COALESCE(SUM(cp.valor_alocado), 0) AS investimento_total
		FROM programa pr
		LEFT JOIN projeto p ON p.id_programa = pr.id_programa
		LEFT JOIN compra_projeto cp ON p.id_projeto = cp.id_projeto
		GROUP BY pr.codigo_programa, pr.nome_programa
		ORDER BY pr.nome_programa`
	
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.ProgramaInvestimento
	for rows.Next() {
		var item entity.ProgramaInvestimento
		err := rows.Scan(
			&item.CodigoPrograma,
			&item.NomePrograma,
			&item.InvestimentoTotal,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	return results, rows.Err()
}

func (r *PostgresProjetoRepository) FindMateriaisByProjeto() ([]entity.ProjetoMaterial, error) {
	query := `
		SELECT
			p.codigo_projeto,
			p.nome_projeto,
			m.codigo_material,
			m.descricao as descricao_material,
			COALESCE(SUM(ep.quantidade), 0) as quantidade_estoque
		FROM projeto p
		INNER JOIN estoque_projeto ep ON p.id_projeto = ep.id_projeto
		INNER JOIN material m ON ep.id_material = m.id_material
		GROUP BY p.codigo_projeto, p.nome_projeto, m.codigo_material, m.descricao
		ORDER BY p.nome_projeto, m.descricao;
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.ProjetoMaterial
	for rows.Next() {
		var item entity.ProjetoMaterial
		err := rows.Scan(
			&item.CodigoProjeto,
			&item.NomeProjeto,
			&item.CodigoMaterial,
			&item.DescricaoMaterial,
			&item.QuantidadeEstoque,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	return results, rows.Err()
}
