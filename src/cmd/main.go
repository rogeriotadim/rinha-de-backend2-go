package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rogeriotadim/rinha-de-backend2-go/cmd/config"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/repository"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/infra/database"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/infra/server/handlers"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/usecase"
)

func main() {
	ctx := context.Background()
	conf, _ := config.LoadConfig("/app")
	dbPool, _ := database.NewDatabasePool(conf, ctx)
	repo := repository.NewClienteRepository(ctx, dbPool)
	addTransacaoUseCase := usecase.NewAddTransacaoUseCase(dbPool, repo)
	getExtratoUseCase := usecase.NewGetExtratoUseCase(dbPool, repo)
	cr := handlers.NewClienteHandler(*addTransacaoUseCase, *getExtratoUseCase)

	r := chi.NewRouter()
	r.Route("/clientes", func(r chi.Router) {
		r.Post("/{id}/transacoes", cr.AddTransacao)
		r.Get("/{id}/extrato", cr.GetExtrato)
	})

	log.Fatal(http.ListenAndServe(":"+conf.WebServerPort, r))

}
