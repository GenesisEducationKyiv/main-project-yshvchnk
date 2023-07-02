package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func RetrieveAndDecodeFromAPI(url string, response interface{}) (float64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("couldn't get a response from the API, status code: %d", resp.StatusCode)
	}

	bodyBuffer := bytes.NewBuffer(body)
	err = json.NewDecoder(bodyBuffer).Decode(response)
	if err != nil {
		return 0, err
	}

	switch r := response.(type) {
		case *CoinGeckoResponse:
			log.Printf("CoinGecko - Response: %s", body)
			return r.Currency.UAH, nil
		case *CryptoCompareResponse:
			log.Printf("CryptoCompare - Response: %s", body)
			return r.UAH, nil
		case *CoinPaprikaResponse:
			log.Printf("CoinPaprika - Response: %s", body)
			return r.UAH, nil
		default:
			return 0, errors.New("unsupported response type")
	}
}