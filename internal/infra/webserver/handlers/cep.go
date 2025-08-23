package handlers

import (
	"net/http"

	"github.com/brecabral/multithreading/internal/services"
	"github.com/brecabral/multithreading/pkg/validators"
	"github.com/go-chi/chi/v5"
)

type CepHandler struct {
	Service *services.AddressService
}

func NewCepHandler(service *services.AddressService) *CepHandler {
	return &CepHandler{
		Service: service,
	}
}

// GetCep godoc
//
//	@Summary		Query CEP
//	@Description	Search CEP concurrently through multiple APIs
//	@Tags			cep
//	@Param			cep	path	string	true	"8-digit CEP without hyphen"
//	@Success		200
//	@Failure		400
//	@Router			/{cep} [get]
func (h *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if !validators.IsValidCep(cep) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.Service.FindAddressByCep(cep)
	w.WriteHeader(http.StatusOK)
}
