package viacep

import "github.com/brecabral/multithreading/internal/domain"

func toAddress(viaCep viaCepResponse) domain.Address {
	return domain.Address{
		Cep:          viaCep.Cep,
		State:        viaCep.UF,
		City:         viaCep.Localidade,
		Neighborhood: viaCep.Bairro,
		Street:       viaCep.Logradouro,
		Api:          "viacep.com.br",
	}
}
