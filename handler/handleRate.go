package handler

import (
	"bitcoin-app/service"
	"encoding/json"
	"net/http"
)

func HandleRate(w http.ResponseWriter, r *http.Request) {

	rate, err := service.GetBitcoinRate()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(rate)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}


