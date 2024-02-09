package model

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/dto"
)

type Cliente struct {
	Id         int   `db:"id"`
	Limite     int64 `db:"limite"`
	Saldo      int64 `db:"saldo"`
	Transacoes []Transacao
}

func (c *Cliente) HydrateDto() dto.ExtratoDTO {
	return dto.ExtratoDTO{
		Saldo:             c.saldoHydrate(),
		UltimasTransacoes: c.ultimasTransacoesHydrate(),
	}
}

func (c *Cliente) saldoHydrate() dto.SaldoDTO {
	return dto.SaldoDTO{
		Limite:      c.Limite,
		Total:       c.Saldo,
		DataExtrato: time.Now(),
	}
}

func (c *Cliente) ultimasTransacoesHydrate() []dto.TransacaoDTO {
	var listaTransacaoDTO []dto.TransacaoDTO = []dto.TransacaoDTO{}
	if len(c.Transacoes) > 0 {
		err := copier.Copy(&listaTransacaoDTO, &c.Transacoes)
		if err != nil {
			return nil
		}
	}
	return listaTransacaoDTO
}
