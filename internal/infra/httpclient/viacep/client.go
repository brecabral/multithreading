package viacep

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/brecabral/multithreading/internal/domain"
	"github.com/brecabral/multithreading/pkg/validators"
)

type ViaCepClient struct {
	Client *http.Client
}

func NewViaCepClient(client *http.Client) *ViaCepClient {
	return &ViaCepClient{
		Client: client,
	}
}

func (c *ViaCepClient) FindAddress(ctx context.Context, cep string) (*domain.Address, error) {
	if !validators.IsValidCep(cep) {
		return nil, errors.New("Invalid CEP")
	}

	url := "http://viacep.com.br/ws/" + cep + "/json/"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response viaCepResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Cep == "" {
		return nil, errors.New("Not Found CEP or API error")
	}

	addr := toAddress(cep, response)
	return addr, nil
}
