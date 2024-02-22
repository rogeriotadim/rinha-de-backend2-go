package model

import (
	"testing"

	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/dto"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransacao(t *testing.T) {
	dtoIn := dto.TransacaoDtoIn{
		Valor:     199000,
		Tipo:      "C",
		Descricao: "Compra de um item no mercado.",
	}

	transacao := CreateTransacao(dtoIn, "1")
	assert.Equal(t, "c", transacao.Tipo)
}
