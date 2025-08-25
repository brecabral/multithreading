package viacep

import "github.com/brecabral/multithreading/internal/domain"

func toAddress(cep string, viaCep viaCepResponse) *domain.Address {
	return &domain.Address{
		Cep:          cep,
		State:        viaCep.UF,
		City:         viaCep.Localidade,
		Neighborhood: viaCep.Bairro,
		Street:       viaCep.Logradouro,
		Api:          "viacep.com.br",
	}
}
