package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

var emailFrom = os.Getenv("EMAIL_FROM")
var emailHost = os.Getenv("EMAIL_HOST")
var emailPort = os.Getenv("EMAIL_PORT")
var emailUsername = os.Getenv("EMAIL_USERNAME")
var emailPassword = os.Getenv("EMAIL_PASSWORD")

func SendEmail(email string, rate float64) error {
	from := emailFrom 
	to := email
	subject := "Bitcoin Rate"
	body := fmt.Sprintf("Bitcoin rate is %.2f UAH", rate)

	smtpHost := emailHost
	smtpPort := emailPort
	smtpUsername := emailUsername
	smtpPassword := emailPassword

	message := []byte(
		"From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n",
	)

	err := smtp.SendMail(smtpHost+":"+smtpPort, smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost), from, []string{to}, message)
	if err != nil {
		log.Printf("Failed to send email to %s: %v\n", email, err)
		return err
	}

	log.Printf("Email sent to %s\n", email)
	return err
}
