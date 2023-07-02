package service

type CoinGeckoResponse struct {
	Currency struct {
		UAH float64 `json:"uah"`
	} `json:"bitcoin"`
}

type CryptoCompareResponse struct {
	UAH float64 `json:"UAH"`
}

type CoinPaprikaResponse struct {
	UAH float64 `json:"price"`
}