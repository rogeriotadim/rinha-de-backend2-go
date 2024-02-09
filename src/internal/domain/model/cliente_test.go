package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHydrateDto(t *testing.T) {
	var cliente Cliente
	cliente.Id = 1
	cliente.Saldo = 100000
	cliente.Limite = 5000000
	var listaTransacoes []Transacao

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 10)
		transacao := Transacao{
			Id:          i,
			ClienteId:   1,
			Valor:       10000,
			Tipo:        "c",
			Descricao:   "teste",
			RealizadaEm: time.Now(),
		}
		listaTransacoes = append(listaTransacoes, transacao)
	}
	cliente.Transacoes = listaTransacoes
	dto := cliente.HydrateDto()
	assert.Equal(t, cliente.Saldo, dto.Saldo.Total)
	assert.Equal(t, cliente.Transacoes[9].Valor, dto.UltimasTransacoes[9].Valor)
}
