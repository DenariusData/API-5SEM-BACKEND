package entity

type DimTempo struct {
	SkTempo      string `json:"sk_tempo"`
	DataCompleta string `json:"data_completa"`
	Ano          string `json:"ano"`
	Mes          string `json:"mes"`
	Dia          string `json:"dia"`
	DiaSemana    string `json:"dia_semana"`
	Semana       string `json:"semana"`
	Semestre     string `json:"semestre"`
	Trimestre    string `json:"trimestre"`
	NomeMes      string `json:"nome_mes"`
}
