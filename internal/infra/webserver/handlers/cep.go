package handlers

import (
	"encoding/json"
	"log"
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

type responseError struct {
	Err string `json:"error"`
}

// GetCep godoc
//
//	@Summary		Query CEP
//	@Description	Search CEP concurrently through multiple APIs
//	@Tags			cep
//	@Param			cep	path		string			true	"8-digit CEP without hyphen"
//
//	@Success		200	{object}	domain.Address	"Returned address from fastest provider"
//	@Failure		400	{object}	responseError	"Invalid CEP"
//	@Failure		504	{object}	responseError	"Timeout or error on APIs"
//
//	@Router			/{cep} [get]
func (h *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cep := chi.URLParam(r, "cep")
	log.Printf("[INFO] request recebida CEP: %v", cep)
	if !validators.IsValidCep(cep) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responseError{Err: "CEP inválido"})
		log.Printf("[ERROR] CEP: %v inválido", cep)
		return
	}

	addr, err := h.Service.FindAddressByCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		json.NewEncoder(w).Encode(responseError{Err: "timeout"})
		log.Printf("[ERROR] timeout: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(addr)
	result, _ := json.MarshalIndent(addr, "", "  ")
	log.Printf("[INFO] Resultado: %v", string(result))
}
