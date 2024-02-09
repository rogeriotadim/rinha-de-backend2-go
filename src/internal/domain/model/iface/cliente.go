package iface

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/model"
)

type ClienteInterface interface {
	GetSaldo(ctx context.Context, tx pgx.Tx, id int) (*model.Cliente, error)
	AddTransacao(ctx context.Context, tx pgx.Tx, transacao *model.Transacao) error
	UpdateSaldo(ctx context.Context, tx pgx.Tx, id int, saldo int64) error
	GetExtrato(id int) (*model.Cliente, error)
}
