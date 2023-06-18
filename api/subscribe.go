package api

import (
	"bitcoin-app/files"
	"fmt"
	"net/http"
)

func HandleSubscribe(w http.ResponseWriter, r *http.Request) {
	//if it post method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//parse form for getting fields
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	//get email field from form
	email := r.Form.Get("email")

	//check if email already subscribed
	if files.IsEmailSubscribed(email) {
		http.Error(w, "Email already subscribed", http.StatusConflict)
		return
	}

	//write email in file
	err = files.SaveEmailToFile(email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	//set status code and provide a response message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Email added")
}