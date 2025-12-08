package api

// MockClient is a mock implementation of the BrasilAPI interface for testing.
type MockClient struct {
	GetCEPFunc        func(cep string) (*CEP, error)
	GetBanksFunc      func() ([]*Bank, error)
	GetBankByCodeFunc func(code string) (*Bank, error)
}

var _ BrasilAPI = (*MockClient)(nil)

func (m *MockClient) GetCEP(cep string) (*CEP, error) {
	if m.GetCEPFunc != nil {
		return m.GetCEPFunc(cep)
	}
	return nil, nil
}

func (m *MockClient) GetBanks() ([]*Bank, error) {
	if m.GetBanksFunc != nil {
		return m.GetBanksFunc()
	}
	return nil, nil
}

func (m *MockClient) GetBankByCode(code string) (*Bank, error) {
	if m.GetBankByCodeFunc != nil {
		return m.GetBankByCodeFunc(code)
	}
	return nil, nil
}
