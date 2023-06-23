package service

import (
	"bitcoin-app/mail"
	"log"
)

func SendEmails(emails []string, rate float64) bool {
	success := true

	for _, email := range emails {
		err := mail.SendEmail(email, rate)
		if err != nil {
			log.Printf("Failed to send email to %s: %v\n", email, err)
		} else {
			success = false
		}
	}

	return success
}