package entity

type DimMaterial struct {
	SkMaterial      string `json:"sk_material"`
	IdMaterial      string `json:"id_material"`
	CodigoMaterial  string `json:"codigo_material"`
	Descricao       string `json:"descricao"`
	Categoria       string `json:"categoria"`
	Fabricante      string `json:"fabricante"`
	CustoEstimado   string `json:"custo_estimado"`
	Status          string `json:"status"`
}
