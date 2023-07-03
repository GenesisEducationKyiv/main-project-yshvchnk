package test

import (
	"bitcoin-app/tests"
	"testing"
)

type MockBitcoinAPI struct{}

func (api *MockBitcoinAPI) GetBitcoinRate() (float64, error) {
	return 1000000.0, nil
}

type MockAPIErrorBitcoinAPI struct{}

func (api *MockAPIErrorBitcoinAPI) GetBitcoinRate() (float64, error) {
	return 0, test.ErrAPIResponse
}

type MockRateNotFoundBitcoinAPI struct{}

func (api *MockRateNotFoundBitcoinAPI) GetBitcoinRate() (float64, error) {
	return 0, test.ErrRateNotFound
}

func TestGetBitcoinRateWillReturnSuccess(t *testing.T) {
	mockAPI := &MockBitcoinAPI{}

	rate, err := mockAPI.GetBitcoinRate()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	expectedRate := 1000000.0
	if rate != expectedRate {
		t.Errorf("Expected rate %f, got: %f", expectedRate, rate)
	}
}

func TestGetBitcoinRateWillReturnAPIError(t *testing.T) {
	mockAPI := &MockAPIErrorBitcoinAPI{}

	rate, err := mockAPI.GetBitcoinRate()

	if err == nil {
		t.Error("Expected an error, got nil")
	}
	expectedError := "couldn't get a response from the API, status code: 500"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got: '%s'", expectedError, err.Error())
	}
	if rate != 0 {
		t.Errorf("Expected rate 0, got: %f", rate)
	}
}

func TestGetBitcoinRateWillReturnRateNotFound(t *testing.T) {
	mockAPI := &MockRateNotFoundBitcoinAPI{}

	rate, err := mockAPI.GetBitcoinRate()

	if err == nil {
		t.Error("Expected an error, got nil")
	}

	expectedError := "failed to retrieve the rate for the bitcoin"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got: '%s'", expectedError, err.Error())
	}

	if rate != 0 {
		t.Errorf("Expected rate 0, got: %f", rate)
	}
}
