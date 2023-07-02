package handler

type CurrencyRateProvider interface {
	GetCurrencyRate() (float64, error)
}

type CurrencyRateService struct {
	Provider CurrencyRateProvider
}

func (s *CurrencyRateService) GetCurrencyRate() (float64, error) {
	return s.Provider.GetCurrencyRate()
}