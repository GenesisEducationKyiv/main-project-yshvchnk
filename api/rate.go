package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

func HandleRate(w http.ResponseWriter, r *http.Request) { //w for sending answear to client, r - for getting details about http
	//if it get method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//getting bitcoin rate
	rate, err := getBitcoinRate()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	//give content-type name
	w.Header().Set("Content-Type", "application/json")

	//encoding bitcoin rate in json
	err = json.NewEncoder(w).Encode(rate)

	//error handling
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

//CoinGecko API
const coinGeckoAPI = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=UAH"

//get current bitcoin rate
func getBitcoinRate() (float64, error) {
	//get request to API
	resp, err := http.Get(coinGeckoAPI)
	if err != nil {
		return 0, err
	}

	// ensures that the response body is closed after it has been read
	defer resp.Body.Close()

	//check if we get a rate
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("invalid status value")
	}

	//create a map variable named data
	var data map[string]map[string]float64

	//decode response into map
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	//get bitcoin rate from data
	rate, ok := data["bitcoin"]["uah"]
	if !ok {
		return 0, errors.New("invalid status value")
	}

	return rate, nil
}