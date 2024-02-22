package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/dto"
)

type Transacao struct {
	Id          int       `db:"id"`
	ClienteId   int       `db:"cliente_id"`
	Valor       int64     `db:"valor"`
	Tipo        string    `db:"tipo"`
	Descricao   string    `db:"descricao"`
	RealizadaEm time.Time `db:"realizada_em"`
}

func CreateTransacao(dto dto.TransacaoDtoIn, id string) Transacao {
	clienteId, _ := strconv.Atoi(id)
	return Transacao{
		ClienteId:   clienteId,
		Valor:       dto.Valor,
		Tipo:        strings.ToLower(dto.Tipo),
		Descricao:   dto.Descricao,
		RealizadaEm: time.Now(),
	}
}
