package main

import (
	"context"
	"log"
	"net/http"

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

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("POST /clientes/{id}/transacoes", cr.AddTransacao)
	mux.HandleFunc("GET /clientes/{id}/extrato", cr.GetExtrato)

	log.Fatal(http.ListenAndServe(":"+conf.WebServerPort, mux))
}
