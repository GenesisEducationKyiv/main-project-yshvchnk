package service

import (
	"fmt"
	"os"
)

var emailFrom = os.Getenv("EMAIL_FROM")

func CreateEmail (to string, rate float64) []byte {
	subject := "Bitcoin Rate"
	body := fmt.Sprintf("Bitcoin rate is %.2f UAH", rate)

	message := []byte(
		"From: " + emailFrom + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n",
	)

	return message
}
