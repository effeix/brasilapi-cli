package tests

import (
	"testing"

	"github.com/effeix/brasilapi-cli/internal/api"
)

func TestNewClient(t *testing.T) {
	client := api.NewClient()

	if client == nil {
		t.Fatal("NewClient returned nil")
	}
}

func TestBrasilAPIError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      api.BrasilAPIError
		expected string
	}{
		{
			name:     "with name",
			err:      api.BrasilAPIError{Name: "NOT_FOUND", Message: "CEP not found"},
			expected: "NOT_FOUND: CEP not found",
		},
		{
			name:     "without name",
			err:      api.BrasilAPIError{Message: "Something went wrong"},
			expected: "Something went wrong",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("BrasilAPIError.Error() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestMockClient_DefaultsToNil(t *testing.T) {
	client := &api.MockClient{}

	cep, err := client.GetCEP("01001000")
	if cep != nil || err != nil {
		t.Errorf("expected nil, nil; got %v, %v", cep, err)
	}

	banks, err := client.GetBanks()
	if banks != nil || err != nil {
		t.Errorf("expected nil, nil; got %v, %v", banks, err)
	}

	bank, err := client.GetBankByCode("1")
	if bank != nil || err != nil {
		t.Errorf("expected nil, nil; got %v, %v", bank, err)
	}
}
