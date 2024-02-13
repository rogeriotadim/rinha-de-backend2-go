package usecase

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/model"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/model/iface"
)

type AddTransacaoUseCase struct {
	dbPool *pgxpool.Pool
	repo   iface.ClienteInterface
}

func NewAddTransacaoUseCase(dbPool *pgxpool.Pool, repo iface.ClienteInterface) *AddTransacaoUseCase {
	return &AddTransacaoUseCase{
		dbPool: dbPool,
		repo:   repo,
	}
}

func (c *AddTransacaoUseCase) Execute(transacao *model.Transacao) (*model.Cliente, error) {
	ctx := context.Background()
	conn, err := c.dbPool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	cliente, err := c.repo.GetSaldo(ctx, tx, transacao.ClienteId)
	if err != nil {
		if err.Error() == "404" {
			return nil, errors.New("o cliente não existe")
		}
		return nil, err
	}

	valor := transacao.Valor
	if ok := transacao.Tipo == "d"; ok {
		valor = valor * -1
	}
	novoSaldo := cliente.Saldo + valor
	if novoSaldo < cliente.Limite*-1 {
		return nil, errors.New("limite")
	}

	err = c.repo.AddTransacao(ctx, tx, transacao)
	if err != nil {
		return nil, err
	}

	err = c.repo.UpdateSaldo(ctx, tx, cliente.Id, novoSaldo)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	cliente.Saldo = novoSaldo
	return cliente, nil
}
