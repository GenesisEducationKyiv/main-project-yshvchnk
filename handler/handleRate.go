package handler

import (
	"bitcoin-app/service"
	"encoding/json"
	"net/http"
)

type BitcoinRateHandler struct {
	BitcoinAPI service.BitcoinAPI
}

func NewBitcoinRateHandler(bitcoinAPI service.BitcoinAPI) *BitcoinRateHandler {
	return &BitcoinRateHandler{
		BitcoinAPI: bitcoinAPI,
	}
}

func (h *BitcoinRateHandler) HandleRate(w http.ResponseWriter, r *http.Request) {

	rate, err := h.BitcoinAPI.GetBitcoinRate()
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


