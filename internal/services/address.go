package services

import (
	"context"
	"time"

	"github.com/brecabral/multithreading/internal/domain"
)

type AddressService struct {
	AddressProviders []domain.Provider
	TimeoutLimit     time.Duration
}

func NewAddressService(addressProviders []domain.Provider, timeoutLimit time.Duration) *AddressService {
	return &AddressService{
		AddressProviders: addressProviders,
		TimeoutLimit:     timeoutLimit,
	}
}

func (s *AddressService) FindAddressByCep(cep string) (*domain.Address, error) {
	resultChan := make(chan *domain.Address, 1)

	ctx, cancel := context.WithTimeout(context.Background(), s.TimeoutLimit)
	defer cancel()

	for _, provider := range s.AddressProviders {
		go queryProvider(ctx, cep, provider, resultChan)
	}

	return awaitFirst(ctx, resultChan)
}

func awaitFirst(ctx context.Context, resultChan <-chan *domain.Address) (*domain.Address, error) {
	select {
	case addr := <-resultChan:
		return addr, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func queryProvider(ctx context.Context, cep string, provider domain.Provider, resultChan chan<- *domain.Address) {
	addr, err := provider.FindAddress(ctx, cep)
	if err != nil {
		return
	}
	select {
	case resultChan <- addr:
		return
	case <-ctx.Done():
		return
	}
}
