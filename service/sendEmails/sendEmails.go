package service

import (
	"log"
)

func (es *EmailSenderDetails) SendEmails(emails []string, rate float64) bool {
	success := true
	
	for _, email := range emails {
		emailText := CreateEmail(email, rate)
		err := SendEmail(emailText, email)
		if err != nil {
			log.Printf("Failed to send email to %s: %v\n", email, err)
		} else {
			success = false
		}
	}

	return success
}