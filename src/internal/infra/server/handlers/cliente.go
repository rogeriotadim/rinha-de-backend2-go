package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/dto"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/domain/model"
	"github.com/rogeriotadim/rinha-de-backend2-go/internal/usecase"
)

type ClienteHandler struct {
	addTransacaoUseCase usecase.AddTransacaoUseCase
	getExtratoUseCase   usecase.GetExtratoUseCase
}

func NewClienteHandler(addTransacaoUseCase usecase.AddTransacaoUseCase, getExtratoUseCase usecase.GetExtratoUseCase) *ClienteHandler {
	return &ClienteHandler{
		addTransacaoUseCase: addTransacaoUseCase,
		getExtratoUseCase:   getExtratoUseCase,
	}
}

func (h *ClienteHandler) AddTransacao(w http.ResponseWriter, r *http.Request) {
	var dtoIn dto.TransacaoDtoIn
	clienteId := r.PathValue("id")

	err := json.NewDecoder(r.Body).Decode(&dtoIn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = dtoIn.Validate()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	transacao := model.CreateTransacao(dtoIn, clienteId)
	cliente, err := h.addTransacaoUseCase.Execute(&transacao)
	if err != nil {
		if err.Error() == "limite" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		if err.Error() == "o cliente não existe" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dtoOut := dto.TransacaoDtoOut{
		Saldo:  cliente.Saldo,
		Limite: cliente.Limite,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dtoOut)
}

func (h *ClienteHandler) GetExtrato(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cliente, err := h.getExtratoUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if cliente == nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}
	dtoOut := cliente.HydrateDto()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dtoOut)
}
