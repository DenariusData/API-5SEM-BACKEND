package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

type PostgresFatoComprasRepository struct {
	db *pgxpool.Pool
}

func NewPostgresFatoComprasRepository(db *pgxpool.Pool) *PostgresFatoComprasRepository {
	return &PostgresFatoComprasRepository{db: db}
}

func (r *PostgresFatoComprasRepository) FindAll() ([]entity.FatoCompras, error) {
	query := `
		SELECT
			pc.id_pedido::text,
			COALESCE(cp.id_projeto::text, ''),
			pc.id_fornecedor::text,
			pc.id_solicitacao::text,
			pc.data_pedido::text,
			COALESCE(pc.valor_total::text, ''),
			COALESCE(cp.valor_alocado::text, '')
		FROM pedido_compra pc
		LEFT JOIN compra_projeto cp ON pc.id_pedido = cp.id_pedido
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.FatoCompras
	for rows.Next() {
		var f entity.FatoCompras
		err := rows.Scan(
			&f.SkFato, &f.SkProjeto, &f.SkFornecedor, &f.SkSolicitacao,
			&f.SkTempo, &f.ValorTotalPedido, &f.ValorAlocadoProjeto,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, f)
	}

	return results, rows.Err()
}
