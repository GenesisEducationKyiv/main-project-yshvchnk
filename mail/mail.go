package mail

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(email string, rate float64) error {
	from := "your-email@example.com" //enter your email
	to := email
	subject := "Bitcoin Rate"
	body := fmt.Sprintf("Bitcoin rate is %.2f UAH", rate)

	//smtp-server set up
	smtpHost := "smtp.example.com" //enter your smtp host
	smtpPort := "587"
	smtpUsername := "username" //enter your username
	smtpPassword := "password" //enter your password

	//message format
	message := []byte(
		"From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n",
	)

	//sending email
	fmt.Printf("Sending email to %s: Bitcoin rate is %.2f UAH\n", email, rate)
	err := smtp.SendMail(smtpHost+":"+smtpPort, smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost), from, []string{to}, message)
	if err != nil {
		log.Printf("Failed to send email to %s: %v\n", email, err)
		return err
	}

	log.Printf("Email sent to %s\n", email)
	return nil
}
