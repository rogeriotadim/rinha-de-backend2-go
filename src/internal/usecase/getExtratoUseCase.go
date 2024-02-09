package usecase

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/model"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/model/iface"
)

type GetExtratoUseCase struct {
	dbPool *pgxpool.Pool
	repo   iface.ClienteInterface
}

func NewGetExtratoUseCase(dbPool *pgxpool.Pool, repo iface.ClienteInterface) *GetExtratoUseCase {
	return &GetExtratoUseCase{dbPool: dbPool, repo: repo}
}

func (g *GetExtratoUseCase) Execute(id int) (*model.Cliente, error) {
	cliente, err := g.repo.GetExtrato(id)
	return cliente, err
}
