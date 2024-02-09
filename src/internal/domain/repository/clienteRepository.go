package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/model"
)

type ClienteRepository struct {
	dbPool *pgxpool.Pool
	ctx    context.Context
}

func NewClienteRepository(ctx context.Context, dbPool *pgxpool.Pool) *ClienteRepository {
	return &ClienteRepository{
		dbPool: dbPool,
		ctx:    ctx,
	}
}

func (r *ClienteRepository) GetSaldo(ctx context.Context, tx pgx.Tx, id int) (*model.Cliente, error) {
	sql := fmt.Sprintf("SELECT id, saldo, limite FROM clientes WHERE id=%d;", id)
	row := tx.QueryRow(ctx, sql)
	var idCliente int
	var saldo, limite int64
	row.Scan(&idCliente, &saldo, &limite)
	if idCliente == 0 {
		return nil, errors.New("404")
	}
	cliente := model.Cliente{
		Id:     idCliente,
		Saldo:  saldo,
		Limite: limite,
	}
	return &cliente, nil
}

func (r *ClienteRepository) AddTransacao(ctx context.Context, tx pgx.Tx, tr *model.Transacao) error {

	layout := "2006-01-02 15:04:05.000"
	timestamp := tr.RealizadaEm.Format(layout)
	sql := fmt.Sprintf(`INSERT INTO transacoes 
		 (cliente_id, valor, tipo, descricao, realizada_em) 
		 VALUES (%d, %d, '%s', '%s', '%v');`,
		tr.ClienteId, tr.Valor, tr.Tipo, tr.Descricao, timestamp)
	_, err := tx.Exec(r.ctx, sql)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClienteRepository) UpdateSaldo(ctx context.Context, tx pgx.Tx, id int, saldo int64) error {
	sql := fmt.Sprintf(`UPDATE clientes 
		 SET saldo=%d WHERE id=%d;`, saldo, id)
	_, err := tx.Exec(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClienteRepository) GetExtrato(id int) (*model.Cliente, error) {
	sql := fmt.Sprintf(`SELECT 
		c.limite, c.saldo, t.valor, t.tipo, t.descricao, t.realizada_em
		FROM clientes c JOIN transacoes t ON c.id = t.cliente_id
		WHERE c.id= %d ORDER BY t.realizada_em DESC LIMIT 10;`, id)
	conn, err := c.dbPool.Acquire(c.ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, _ := conn.Query(c.ctx, sql)
	defer rows.Close()
	cliente := model.Cliente{}
	transacoes := []model.Transacao{}
	for rows.Next() {
		t := model.Transacao{}
		err := rows.Scan(&cliente.Limite, &cliente.Saldo,
			&t.Valor, &t.Tipo, &t.Descricao, &t.RealizadaEm)
		if err != nil {
			return nil, err
		}
		transacoes = append(transacoes, t)
	}
	cliente.Transacoes = transacoes
	return &cliente, nil
}
