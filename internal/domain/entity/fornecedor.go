package entity

type DimFornecedor struct {
	SkFornecedor     string `json:"sk_fornecedor"`
	IdFornecedor     string `json:"id_fornecedor"`
	CodigoFornecedor string `json:"codigo_fornecedor"`
	RazaoSocial      string `json:"razao_social"`
	Cidade           string `json:"cidade"`
	Estado           string `json:"estado"`
	Categoria        string `json:"categoria"`
	Status           string `json:"status"`
}
