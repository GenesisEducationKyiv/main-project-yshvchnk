package handler

import (
	"bitcoin-app/service/getCurrencyRate"
)

type CurrencyRateHandler struct {
	CurrencyAPI service.CurrencyRateProvider
}

func NewCurrencyRateHandler(currencyAPI service.CurrencyRateProvider) *CurrencyRateHandler {
	return &CurrencyRateHandler{
		CurrencyAPI: currencyAPI,
	}
}