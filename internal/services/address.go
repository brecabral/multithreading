package services

import (
	"log"
	"time"

	"github.com/brecabral/multithreading/internal/domain"
)

type AddressService struct {
	AddressProviders []domain.Provider
	result           chan domain.Address
}

func NewAddressService(addressProviders []domain.Provider) *AddressService {
	return &AddressService{
		AddressProviders: addressProviders,
		result:           make(chan domain.Address),
	}
}

func (s *AddressService) FindAddressByCep(cep string) {
	for _, addressProvider := range s.AddressProviders {
		go func(provider domain.Provider) {
			addr, _ := provider.FindAddress(cep)
			s.result <- addr
		}(addressProvider)
	}
	s.printResult()
}

func (s *AddressService) printResult() {
	select {
	case addr := <-s.result:
		log.Printf("[RESULTADO] %v", addr)
	case <-time.After(1 * time.Second):
		log.Print("[ERROR] timeout")
	}
}
