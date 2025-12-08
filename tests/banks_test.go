package tests

import (
	"errors"
	"testing"

	"github.com/effeix/brasilapi-cli/internal/api"
)

func TestGetBanks_Success(t *testing.T) {
	expected := []*api.Bank{
		{ISPB: "00000000", Name: "BCB", Code: 1, FullName: "Banco Central do Brasil"},
		{ISPB: "00000208", Name: "BRB", Code: 70, FullName: "BRB - BANCO DE BRASILIA S.A."},
	}

	client := &api.MockClient{
		GetBanksFunc: func() ([]*api.Bank, error) {
			return expected, nil
		},
	}

	result, err := client.GetBanks()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d banks, got %d", len(expected), len(result))
	}

	if result[0].Code != expected[0].Code {
		t.Errorf("expected code %d, got %d", expected[0].Code, result[0].Code)
	}

	if result[0].ISPB != expected[0].ISPB {
		t.Errorf("expected ISPB %s, got %s", expected[0].ISPB, result[0].ISPB)
	}
}

func TestGetBanks_Empty(t *testing.T) {
	client := &api.MockClient{
		GetBanksFunc: func() ([]*api.Bank, error) {
			return []*api.Bank{}, nil
		},
	}

	result, err := client.GetBanks()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("expected 0 banks, got %d", len(result))
	}
}

func TestGetBanks_Error(t *testing.T) {
	networkErr := errors.New("network connection failed")

	client := &api.MockClient{
		GetBanksFunc: func() ([]*api.Bank, error) {
			return nil, networkErr
		},
	}

	result, err := client.GetBanks()

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
}

func TestGetBankByCode_Success(t *testing.T) {
	expected := &api.Bank{
		ISPB:     "00000000",
		Name:     "BCB",
		Code:     1,
		FullName: "Banco Central do Brasil",
	}

	client := &api.MockClient{
		GetBankByCodeFunc: func(code string) (*api.Bank, error) {
			if code != "1" {
				t.Errorf("unexpected code: %s", code)
			}
			return expected, nil
		},
	}

	result, err := client.GetBankByCode("1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Code != expected.Code {
		t.Errorf("expected code %d, got %d", expected.Code, result.Code)
	}

	if result.Name != expected.Name {
		t.Errorf("expected name %s, got %s", expected.Name, result.Name)
	}

	if result.FullName != expected.FullName {
		t.Errorf("expected fullName %s, got %s", expected.FullName, result.FullName)
	}
}

func TestGetBankByCode_NotFound(t *testing.T) {
	expectedErr := &api.BrasilAPIError{
		Message: "Código bancário não encontrado",
		Type:    "BANK_CODE_NOT_FOUND",
	}

	client := &api.MockClient{
		GetBankByCodeFunc: func(code string) (*api.Bank, error) {
			return nil, expectedErr
		},
	}

	result, err := client.GetBankByCode("99999")

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

	if apiErr.Type != "BANK_CODE_NOT_FOUND" {
		t.Errorf("expected type 'BANK_CODE_NOT_FOUND', got %q", apiErr.Type)
	}
}

func TestGetBankByCode_NetworkError(t *testing.T) {
	networkErr := errors.New("network connection failed")

	client := &api.MockClient{
		GetBankByCodeFunc: func(code string) (*api.Bank, error) {
			return nil, networkErr
		},
	}

	result, err := client.GetBankByCode("1")

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
