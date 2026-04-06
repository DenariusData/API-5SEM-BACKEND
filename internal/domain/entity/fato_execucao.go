package entity

type FatoExecucaoTarefas struct {
	SkFato          string `json:"sk_fato"`
	SkProjeto       string `json:"sk_projeto"`
	SkTarefa        string `json:"sk_tarefa"`
	SkResponsavel   string `json:"sk_responsavel"`
	SkTempo         string `json:"sk_tempo"`
	HorasTrabalhadas string `json:"horas_trabalhadas"`
}
