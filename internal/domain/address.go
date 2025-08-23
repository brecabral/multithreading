package domain

type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Api          string `json:"api"`
}

type Provider interface {
	FindAddress(cep string) (Address, error)
}
