package service

import (
	"errors"
)

func NewAPIProvider(providers []ExchangeRateProvider) *APIProvider {
	return &APIProvider{
		providers: providers,
	}
}

func (api *APIProvider) GetCurrencyRate() (float64, error) {

	for _, provider := range api.providers {
		rate, err := api.getCurrencyRateFromProvider(provider)
		if err == nil {
			return rate, nil
		}
	}

	return 0, errors.New("unable to retrieve currency rate from any provider")
}

func (api *APIProvider) getCurrencyRateFromProvider(provider ExchangeRateProvider) (float64, error) {
	rate, err := provider.GetCurrencyRate()
	if err != nil {
		return 0, err
	}

	return rate, nil
}