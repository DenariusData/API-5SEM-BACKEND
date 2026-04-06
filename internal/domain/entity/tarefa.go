package entity

type DimTarefa struct {
	SkTarefa        string `json:"sk_tarefa"`
	IdTarefa        string `json:"id_tarefa"`
	CodigoTarefa    string `json:"codigo_tarefa"`
	Titulo          string `json:"titulo"`
	Status          string `json:"status"`
	EstimativaHoras string `json:"estimativa_horas"`
	DataInicio      string `json:"data_inicio"`
	DataFimPrevista string `json:"data_fim_prevista"`
}
