package service

import (
	"log"
	"net/smtp"
	"os"
)

var (
	emailHost = os.Getenv("EMAIL_HOST")
	emailPort = os.Getenv("EMAIL_PORT")
	emailUsername = os.Getenv("EMAIL_USERNAME")
	emailPassword = os.Getenv("EMAIL_PASSWORD")
)

func SendEmail(emailText []byte, to string) error {
	smtpHost := emailHost
	smtpPort := emailPort
	smtpUsername := emailUsername
	smtpPassword := emailPassword

	err := smtp.SendMail(smtpHost+":"+smtpPort, smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost), emailFrom, []string{to}, emailText)
	if err != nil {
		log.Printf("Failed to send email to %s: %v\n", emailText, err)
		return err
	}

	log.Printf("Email sent to %s\n", emailText)
	return err
}
