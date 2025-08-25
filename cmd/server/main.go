package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/brecabral/multithreading/docs"
	"github.com/brecabral/multithreading/internal/domain"
	"github.com/brecabral/multithreading/internal/infra/httpclient/brasilapi"
	"github.com/brecabral/multithreading/internal/infra/httpclient/viacep"
	"github.com/brecabral/multithreading/internal/infra/webserver/handlers"
	"github.com/brecabral/multithreading/internal/services"
	"github.com/go-chi/chi/v5"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

const TIMEOUT_LIMIT = 1 * time.Second

//	@title			CEP Multithreading API
//	@version		1.0
//	@description	Query multiple providers concurrently and return the first response.

// @host		localhost:8000
// @BasePath	/
func main() {
	httpClient := &http.Client{}

	viaCepClient := viacep.NewViaCepClient(httpClient)
	brasilApiClient := brasilapi.NewBrasilApiClient(httpClient)
	providers := []domain.Provider{viaCepClient, brasilApiClient}

	addressService := services.NewAddressService(providers, TIMEOUT_LIMIT)

	cepHandler := handlers.NewCepHandler(addressService)

	router := chi.NewRouter()
	router.Get("/{cep}", cepHandler.GetCep)
	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	log.Fatal(http.ListenAndServe(":8000", router))
}
