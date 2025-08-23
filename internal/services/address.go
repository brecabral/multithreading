package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/brecabral/multithreading/internal/domain"
)

type AddressService struct {
	AddressProviders []domain.Provider
}

func NewAddressService(addressProviders []domain.Provider) *AddressService {
	return &AddressService{
		AddressProviders: addressProviders,
	}
}

func (s *AddressService) FindAddressByCep(cep string) {
	resultChan := make(chan domain.Address)
	defer close(resultChan)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, addressProvider := range s.AddressProviders {
		go func(provider domain.Provider) {
			addr, err := provider.FindAddress(ctx, cep)
			if err == nil {
				resultChan <- addr
			}
		}(addressProvider)
	}

	select {
	case addr := <-resultChan:
		jsonAddr, _ := json.MarshalIndent(addr, "", "  ")
		log.Printf("[RESULTADO]\n%v", string(jsonAddr))
	case <-time.After(1 * time.Second):
		log.Print("[ERROR] timeout")
	}
}
