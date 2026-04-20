package entity

type ProgramaInvestimento struct {
	CodigoPrograma string `json:"codigo_programa"`
	NomePrograma   string `json:"nome_programa"`
	InvestimentoTotal float64 `json:"investimento_total"`
}