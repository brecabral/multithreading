package domain

import "context"

type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Api          string `json:"api"`
}

type Provider interface {
	FindAddress(ctx context.Context, cep string) (*Address, error)
}
