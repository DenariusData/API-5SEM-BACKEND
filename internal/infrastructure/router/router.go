package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/adapter/handler"
)

type Handlers struct {
	Projeto      *handler.ProjetoHandler
	Fornecedor   *handler.FornecedorHandler
	Material     *handler.MaterialHandler
	Responsavel  *handler.ResponsavelHandler
	Solicitacao  *handler.SolicitacaoHandler
	Tarefa       *handler.TarefaHandler
	Tempo        *handler.TempoHandler
	FatoCompras  *handler.FatoComprasHandler
	FatoEstoque  *handler.FatoEstoqueHandler
	FatoExecucao *handler.FatoExecucaoHandler
}

func NewRouter(h Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Route("/dim", func(r chi.Router) {
			r.Get("/projetos", h.Projeto.GetAll)
			r.Get("/fornecedores", h.Fornecedor.GetAll)
			r.Get("/materiais", h.Material.GetAll)
			r.Get("/responsaveis", h.Responsavel.GetAll)
			r.Get("/solicitacoes", h.Solicitacao.GetAll)
			r.Get("/tarefas", h.Tarefa.GetAll)
			r.Get("/tempo", h.Tempo.GetAll)
		})

		r.Route("/fato", func(r chi.Router) {
			r.Get("/compras", h.FatoCompras.GetAll)
			r.Get("/estoque-materiais", h.FatoEstoque.GetAll)
			r.Get("/execucao-tarefas", h.FatoExecucao.GetAll)
		})
	})

	return r
}
