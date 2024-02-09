package dto

import (
	"time"
)

type ExtratoDTO struct {
	Saldo             SaldoDTO       `json:"saldo"`
	UltimasTransacoes []TransacaoDTO `json:"ultimas_transacoes"`
}

type SaldoDTO struct {
	Limite      int64     `json:"limite"`
	Total       int64     `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
}

type TransacaoDTO struct {
	Valor       int64     `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}
