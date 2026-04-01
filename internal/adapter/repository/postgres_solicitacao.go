package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

type PostgresSolicitacaoRepository struct {
	db *pgxpool.Pool
}

func NewPostgresSolicitacaoRepository(db *pgxpool.Pool) *PostgresSolicitacaoRepository {
	return &PostgresSolicitacaoRepository{db: db}
}

func (r *PostgresSolicitacaoRepository) FindAll() ([]entity.DimSolicitacao, error) {
	query := `
		SELECT
			sc.id_solicitacao::text,
			sc.id_solicitacao::text,
			sc.numero_solicitacao,
			COALESCE(sc.prioridade, ''),
			COALESCE(sc.status, ''),
			COALESCE(pc.numero_pedido, ''),
			COALESCE(pc.status, ''),
			COALESCE(pc.data_previsao_entrega::text, '')
		FROM solicitacao_compra sc
		LEFT JOIN pedido_compra pc ON sc.id_solicitacao = pc.id_solicitacao
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DimSolicitacao
	for rows.Next() {
		var s entity.DimSolicitacao
		err := rows.Scan(
			&s.SkSolicitacao, &s.IdSolicitacao, &s.NumeroSolicitacao, &s.Prioridade,
			&s.StatusSolicitacao, &s.NumeroPedido, &s.StatusPedido, &s.DataPrevisaoEntrega,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	return results, rows.Err()
}
