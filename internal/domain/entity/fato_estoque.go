package entity

type FatoEstoqueMateriais struct {
	SkFato               string `json:"sk_fato"`
	SkProjeto            string `json:"sk_projeto"`
	SkMaterial           string `json:"sk_material"`
	SkTempo              string `json:"sk_tempo"`
	QuantidadeEstoque    string `json:"quantidade_estoque"`
	QuantidadeEmpenhada  string `json:"quantidade_empenhada"`
}
