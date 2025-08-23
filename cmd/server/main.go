package main

import (
	"log"
	"net/http"

	"github.com/brecabral/multithreading/internal/domain"
	"github.com/brecabral/multithreading/internal/infra/httpclient/brasilapi"
	"github.com/brecabral/multithreading/internal/infra/httpclient/viacep"
	"github.com/brecabral/multithreading/internal/infra/webserver/handlers"
	"github.com/brecabral/multithreading/internal/services"
	"github.com/go-chi/chi/v5"
)

func main() {
	httpClient := &http.Client{}

	viaCepClient := viacep.NewViaCepClient(httpClient)
	brasilApiClient := brasilapi.NewBrasilApiClient(httpClient)
	providers := []domain.Provider{viaCepClient, brasilApiClient}

	addressService := services.NewAddressService(providers)

	cepHandler := handlers.NewCepHandler(addressService)

	router := chi.NewRouter()
	router.Get("/{cep}", cepHandler.GetCep)

	log.Fatal(http.ListenAndServe(":8000", router))
}
