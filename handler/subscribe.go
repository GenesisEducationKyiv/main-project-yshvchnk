package handler

import (
	"bitcoin-app/service"
	"fmt"
	"net/http"
)

func HandleSubscribe(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")

	err = service.SubscribeEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Email added")
}

