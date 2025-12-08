package api

import "fmt"

type CEP struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (c *Client) GetCEP(cep string) (*CEP, error) {
	endpoint := fmt.Sprintf("/cep/v1/%s", cep)

	var result *CEP
	if err := c.doRequest(endpoint, &result); err != nil {
		return nil, err
	}

	return result, nil
}
