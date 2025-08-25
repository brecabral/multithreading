package brasilapi

import "github.com/brecabral/multithreading/internal/domain"

func toAddress(cep string, brasilApi brasilApiResponse) *domain.Address {
	return &domain.Address{
		Cep:          cep,
		State:        brasilApi.State,
		City:         brasilApi.City,
		Neighborhood: brasilApi.Neighborhood,
		Street:       brasilApi.Street,
		Api:          "brasilapi.com.br",
	}
}
