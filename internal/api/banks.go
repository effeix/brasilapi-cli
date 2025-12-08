package api

import "fmt"

type Bank struct {
	ISPB     string `json:"ispb"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Code     int    `json:"code"`
}

func (c *Client) GetBanks() ([]*Bank, error) {
	endpoint := "/banks/v1"

	var result []*Bank
	if err := c.doRequest(endpoint, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetBankByCode(code string) (*Bank, error) {
	endpoint := fmt.Sprintf("/banks/v1/%s", code)

	var result Bank
	if err := c.doRequest(endpoint, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
