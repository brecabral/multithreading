package brasilapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/brecabral/multithreading/internal/domain"
	"github.com/brecabral/multithreading/pkg/validators"
)

type BrasilApiClient struct {
	Client http.Client
}

func NewBrasilApiClient(client http.Client) *BrasilApiClient {
	return &BrasilApiClient{
		Client: client,
	}
}

func (c *BrasilApiClient) FindAddress(cep string) (domain.Address, error) {
	var addr domain.Address

	if !validators.IsValidCep(cep) {
		return addr, errors.New("Invalid CEP")
	}

	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return addr, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return addr, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return addr, err
	}

	var response brasilApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return addr, err
	}

	addr = toAddress(response)
	return addr, nil
}
