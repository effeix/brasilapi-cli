package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	BaseURL        = "https://brasilapi.com.br/api"
	DefaultTimeout = 30 * time.Second
)

// BrasilAPI defines the interface for interacting with the BrasilAPI service.
// All clients and mock clients must implement this interface.
type BrasilAPI interface {
	GetCEP(cep string) (*CEP, error)
	GetBanks() ([]*Bank, error)
	GetBankByCode(code string) (*Bank, error)
}

type Client struct {
	httpClient *http.Client
	baseURL    string
}

var _ BrasilAPI = (*Client)(nil)

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		baseURL: BaseURL,
	}
}

func (c *Client) doRequest(endpoint string, result any) error {
	url := c.baseURL + endpoint

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr BrasilAPIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
		}

		return &apiErr
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

type BrasilAPIErrorErrors struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Service string `json:"service"`
}

type BrasilAPIError struct {
	Name    string                 `json:"name"`
	Message string                 `json:"message"`
	Type    string                 `json:"type"`
	Errors  []BrasilAPIErrorErrors `json:"errors,omitempty"`
}

func (e *BrasilAPIError) Error() string {
	if e.Name != "" {
		return fmt.Sprintf("%s: %s", e.Name, e.Message)
	}

	return e.Message
}
