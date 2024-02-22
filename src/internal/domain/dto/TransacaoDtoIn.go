package dto

import "errors"

type TransacaoDtoIn struct {
	Valor     int64  `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

func (t *TransacaoDtoIn) Validate() error {
	erroValidacao := errors.New("erro de validação")
	if t.Tipo == "" || t.Descricao == "" || t.Valor <= 0 {
		return erroValidacao
	}

	if t.Tipo != "c" && t.Tipo != "d" {
		return erroValidacao
	}
	return nil
}
