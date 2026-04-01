package entity

type DimProjeto struct {
	SkProjeto       string `json:"sk_projeto"`
	IdProjeto       string `json:"id_projeto"`
	CodigoProjeto   string `json:"codigo_projeto"`
	NomeProjeto     string `json:"nome_projeto"`
	Responsavel     string `json:"responsavel"`
	Status          string `json:"status"`
	CodigoPrograma  string `json:"codigo_programa"`
	NomePrograma    string `json:"nome_programa"`
	GerentePrograma string `json:"gerente_programa"`
	DataInicio      string `json:"data_inicio"`
	DataFimPrevista string `json:"data_fim_prevista"`
}
