package service

import (
	"bitcoin-app/utils"
)

func (api *APIProvider) GetCurrencyRate() (float64, error) {
	rate, err := api.coinGeckoProvider.GetCurrencyRateFromCoinGecko()

	if err != nil {
		rate, err = api.cryptoCompareProvider.GetCurrencyRateFromCryptoCompare()
		if err != nil {
			rate, err = api.coinPaprikaProvider.GetCurrencyRateFromCoinPaprika()
			if err != nil {
				return 0, err
			}
		}
	}

	return rate, nil
}

func (cg *CoinGeckoAPI) GetCurrencyRateFromCoinGecko() (float64, error) {
	utils.LoadEnv()

	url := cg.CoinGeckoURL + "?ids=" + cg.CoinGeckoIds + "&vs_currencies=" + cg.CoinGeckoVsCurrencies

	var coinGeckoResp CoinGeckoResponse

	return RetrieveAndDecodeFromAPI(url, &coinGeckoResp)
}

func (cc *CryptoCompareAPI) GetCurrencyRateFromCryptoCompare() (float64, error) {
	utils.LoadEnv()

	url := cc.CryptoCompareURL + "?fsym=" + cc.CryptoCompareFrom + "&tsyms=" + cc.CryptoCompareTo

	var cryptoCompareResp CryptoCompareResponse

	return RetrieveAndDecodeFromAPI(url, &cryptoCompareResp)
}

func (cp *CoinPaprikaAPI) GetCurrencyRateFromCoinPaprika() (float64, error) {
	utils.LoadEnv()

	url := cp.CoinPaprikaURL + "?base_currency_id=" + cp.CoinPaprikaFrom + "&quote_currency_id=" + cp.CoinPaprikaTo + "&amount=" + cp.CoinPaprikaAmount

	var coinPaprikaResp CoinPaprikaResponse

	return RetrieveAndDecodeFromAPI(url, &coinPaprikaResp)
}