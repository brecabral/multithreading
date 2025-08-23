package brasilapi

import "github.com/brecabral/multithreading/internal/domain"

func toAddress(brasilApi brasilApiResponse) domain.Address {
	return domain.Address{
		Cep:          brasilApi.Cep,
		State:        brasilApi.State,
		City:         brasilApi.City,
		Neighborhood: brasilApi.Neighborhood,
		Street:       brasilApi.Street,
		Api:          "brasilapi.com.br",
	}
}
