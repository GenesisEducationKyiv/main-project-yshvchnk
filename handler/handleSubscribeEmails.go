package handler

import (
	"bitcoin-app/service"
	"errors"
	"fmt"
	"net/http"
)


func HandleSubscribe (emailService *service.EmailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")

	// emailService := service.NewEmailService()

	err = emailService.SubscribeEmail(email)
	if err != nil {
		if errors.Is(err, service.ErrEmailAlreadySubscribed) {
			http.Error(w, "Email is already subscribed", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to subscribe email", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Email added")
	}
}

