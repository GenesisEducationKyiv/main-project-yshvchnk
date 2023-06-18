package api

import (
	"bitcoin-app/files"
	"bitcoin-app/mail"
	"fmt"
	"log"
	"net/http"
)

func HandleSendEmails(w http.ResponseWriter, r *http.Request) {
	//if it post method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//getting bitcoin rate
	rate, err := getBitcoinRate()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	//get emails from file
	emails, err := files.GetEmailsFromFile()
	if err != nil {
		http.Error(w, "Failed to load email addresses", http.StatusInternalServerError)
		return
	}

	//send info to subscribed emails
	for _, email := range emails {
		err := mail.SendEmail(email, rate)
		if err != nil {
			log.Printf("Failed to send email to %s: %v\n", email, err)
		}
	}

	//set status code and provide a response message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Emails have been sent")
}