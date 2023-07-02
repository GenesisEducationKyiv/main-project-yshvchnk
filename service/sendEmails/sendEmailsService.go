package service

import (
	"bitcoin-app/store"
	"fmt"
	"github.com/pkg/errors"
)

type EmailSender interface {
	SendEmails(emails []string, rate float64) bool
	GetCurrencyRate() (float64, error)
}

type EmailSendService struct {
	Storage store.EmailStorage
	Sender  EmailSender
}

func (s *EmailSendService) SendEmails() error {
	emails, err := s.Storage.GetEmailsFromFile()
	if err != nil {
		return errors.Wrap(err, "Failed to load email addresses")
	}

	rate, err := s.Sender.GetCurrencyRate()
	if err != nil {
		return errors.Wrap(err, "Failed to get Bitcoin rate")
	}

	success := s.Sender.SendEmails(emails, rate)
	if !success {
		return fmt.Errorf("Failed to send %d emails", len(emails))
	}

	return nil
}
