package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type CurrencyProviderResponse interface {
	GetCurrencyRate() float64
}

type LoggingResponse struct {
	Response CurrencyProviderResponse
}

func (lr *LoggingResponse) GetCurrencyRate() float64 {
	responseValue := lr.Response.GetCurrencyRate()
	formattedResponse := strconv.FormatFloat(responseValue, 'f', -1, 64)
	log.Printf("Response: &{%s}", formattedResponse)
	return responseValue
}

func RetrieveAndDecodeFromAPI(url string, response CurrencyProviderResponse) (float64, error) {
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

	loggedResponse := &LoggingResponse{Response: response}
	return loggedResponse.GetCurrencyRate(), nil
}