package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

type BitcoinAPI interface {
	GetBitcoinRate() (float64, error)
}

type CoinGeckoAPI struct {
	CoinGeckoURL string
	Ids          string
	VsCurrencies string
}

func NewCoinGeckoAPI() *CoinGeckoAPI {
	return &CoinGeckoAPI{
		CoinGeckoURL: os.Getenv("COIN_GECKO_API"),
		Ids:          os.Getenv("IDS"),
		VsCurrencies: os.Getenv("VS_CURRENCIES"),
	}
}

type BitcoinResponse struct {
	Bitcoin struct {
		UAH float64 `json:"uah"`
	} `json:"bitcoin"`
}

func (api *CoinGeckoAPI) GetBitcoinRate() (float64, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := api.CoinGeckoURL + "?ids=" + api.Ids + "&vs_currencies=" + api.VsCurrencies

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("couldn't get a response from the API, status code: %d", resp.StatusCode)
	}

	var data BitcoinResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	rate := data.Bitcoin.UAH
	if rate == 0 {
		return 0, errors.New("failed to retrieve the rate for the bitcoin")
	}

	return rate, nil
}