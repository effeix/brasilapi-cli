package tests

import (
	"errors"
	"testing"

	"github.com/effeix/brasilapi-cli/internal/api"
)

func TestGetCEP_Success(t *testing.T) {
	expected := &api.CEP{
		CEP:          "01001000",
		State:        "SP",
		City:         "São Paulo",
		Neighborhood: "Sé",
		Street:       "Praça da Sé",
		Service:      "open-cep",
	}

	client := &api.MockClient{
		GetCEPFunc: func(cep string) (*api.CEP, error) {
			if cep != expected.CEP {
				t.Errorf("unexpected CEP: %s", cep)
			}
			return expected, nil
		},
	}

	result, err := client.GetCEP(expected.CEP)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.CEP != expected.CEP {
		t.Errorf("expected CEP %s, got %s", expected.CEP, result.CEP)
	}

	if result.State != expected.State {
		t.Errorf("expected State %s, got %s", expected.State, result.State)
	}

	if result.City != expected.City {
		t.Errorf("expected City %s, got %s", expected.City, result.City)
	}

	if result.Neighborhood != expected.Neighborhood {
		t.Errorf("expected Neighborhood %s, got %s", expected.Neighborhood, result.Neighborhood)
	}

	if result.Street != expected.Street {
		t.Errorf("expected Street %s, got %s", expected.Street, result.Street)
	}
}

func TestGetCEP_NotFound(t *testing.T) {
	expectedErr := &api.BrasilAPIError{
		Name:    "CepPromiseError",
		Message: "Todos os serviços de CEP retornaram erro.",
		Type:    "service_error",
	}

	client := &api.MockClient{
		GetCEPFunc: func(cep string) (*api.CEP, error) {
			return nil, expectedErr
		},
	}

	result, err := client.GetCEP("00000000")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}

	apiErr, ok := err.(*api.BrasilAPIError)
	if !ok {
		t.Fatalf("expected *api.BrasilAPIError, got %T", err)
	}

	if apiErr.Type != "service_error" {
		t.Errorf("expected type 'service_error', got %q", apiErr.Type)
	}
}

func TestGetCEP_InvalidCEP(t *testing.T) {
	expectedErr := &api.BrasilAPIError{
		Name:    "CepPromiseError",
		Message: "CEP deve conter exatamente 8 caracteres.",
		Type:    "validation_error",
	}

	client := &api.MockClient{
		GetCEPFunc: func(cep string) (*api.CEP, error) {
			return nil, expectedErr
		},
	}

	result, err := client.GetCEP("123")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}

	apiErr, ok := err.(*api.BrasilAPIError)
	if !ok {
		t.Fatalf("expected *api.BrasilAPIError, got %T", err)
	}

	if apiErr.Type != "validation_error" {
		t.Errorf("expected type 'validation_error', got %q", apiErr.Type)
	}
}

func TestGetCEP_NetworkError(t *testing.T) {
	networkErr := errors.New("network connection failed")

	client := &api.MockClient{
		GetCEPFunc: func(cep string) (*api.CEP, error) {
			return nil, networkErr
		},
	}

	result, err := client.GetCEP("01001000")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}

	if err.Error() != "network connection failed" {
		t.Errorf("expected network error message, got %q", err.Error())
	}
}
