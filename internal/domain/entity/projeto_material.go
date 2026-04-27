package entity

type ProjetoMaterial struct {
	CodigoProjeto      string `json:"codigo_projeto"`
	NomeProjeto        string `json:"nome_projeto"`
	CodigoMaterial     string `json:"codigo_material"`
	DescricaoMaterial  string `json:"descricao_material"`
	QuantidadeEstoque  int    `json:"quantidade_estoque"`
}
