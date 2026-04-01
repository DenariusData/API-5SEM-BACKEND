package entity

type DimSolicitacao struct {
	SkSolicitacao        string `json:"sk_solicitacao"`
	IdSolicitacao        string `json:"id_solicitacao"`
	NumeroSolicitacao    string `json:"numero_solicitacao"`
	Prioridade           string `json:"prioridade"`
	StatusSolicitacao    string `json:"status_solicitacao"`
	NumeroPedido         string `json:"numero_pedido"`
	StatusPedido         string `json:"status_pedido"`
	DataPrevisaoEntrega  string `json:"data_previsao_entrega"`
}
