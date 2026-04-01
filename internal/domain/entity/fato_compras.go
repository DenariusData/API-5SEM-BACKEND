package entity

type FatoCompras struct {
	SkFato               string `json:"sk_fato"`
	SkProjeto            string `json:"sk_projeto"`
	SkFornecedor         string `json:"sk_fornecedor"`
	SkSolicitacao        string `json:"sk_solicitacao"`
	SkTempo              string `json:"sk_tempo"`
	ValorTotalPedido     string `json:"valor_total_pedido"`
	ValorAlocadoProjeto  string `json:"valor_alocado_projeto"`
}
