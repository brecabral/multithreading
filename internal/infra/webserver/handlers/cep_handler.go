package handlers

import (
	"net/http"

	"github.com/brecabral/multithreading/pkg/validators"
	"github.com/go-chi/chi/v5"
)

type CepHandler struct{}

func (h *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if !validators.IsValidCep(cep) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
