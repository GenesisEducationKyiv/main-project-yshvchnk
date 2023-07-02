package service

import (
	"os"
)

type CurrencyRateProvider interface {
	GetCurrencyRate() (float64, error)
}

type APIProvider struct {
	coinGeckoProvider CoinGeckoAPI
	cryptoCompareProvider CryptoCompareAPI
	coinPaprikaProvider CoinPaprikaAPI
}

func NewCurrencyAPIProvider() *APIProvider {
	coinGeckoProvider := CoinGeckoAPI{
		CoinGeckoURL: os.Getenv("COIN_GECKO_API"),
		CoinGeckoIds: os.Getenv("COIN_GECKO_IDS"),
		CoinGeckoVsCurrencies: os.Getenv("COIN_GECKO_VS_CURRENCIES"),
	}

	cryptoCompareProvider := CryptoCompareAPI{
		CryptoCompareURL: os.Getenv("CRYPTO_COMPARE_URL"),
		CryptoCompareFrom: os.Getenv("CRYPTO_COMPARE_FROM"),
		CryptoCompareTo: os.Getenv("CRYPTO_COMPARE_TO"),
	}

	coinPaprikaProvider := CoinPaprikaAPI{
		CoinPaprikaURL: os.Getenv("COIN_PAPRIKA_URL"),
		CoinPaprikaFrom: os.Getenv("COIN_PAPRIKA_FROM"),
		CoinPaprikaTo: os.Getenv("COIN_PAPRIKA_TO"),
		CoinPaprikaAmount: os.Getenv("COIN_PAPRIKA_AMOUNT"),
	}

	return &APIProvider{
		coinGeckoProvider:  coinGeckoProvider,
		cryptoCompareProvider:  cryptoCompareProvider,
		coinPaprikaProvider:  coinPaprikaProvider,
	}
}