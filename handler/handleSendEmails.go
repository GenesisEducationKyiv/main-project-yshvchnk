package handler

import (
	"bitcoin-app/store"
	"fmt"
	"net/http"
	"github.com/pkg/errors"
)

type BitcoinRateProvider interface {
	GetBitcoinRate() (float64, error)
}

type BitcoinRateService struct {
	Provider BitcoinRateProvider
}

func (s *BitcoinRateService) GetRate() (float64, error) {
	return s.Provider.GetBitcoinRate()
}

type EmailSender interface {
	SendRateToUsers(emails []string, rate float64) bool
}

type EmailService struct {
	Storage store.EmailStorage
	Sender  EmailSender
	RateProvider BitcoinRateProvider
}

func (s *EmailService) SendEmails() error {
	emails, err := s.Storage.GetEmailsFromFile()
	if err != nil {
		return errors.Wrap(err, "Failed to load email addresses")
	}

	rate, err := s.RateProvider.GetBitcoinRate()
	if err != nil {
		return errors.Wrap(err, "Failed to get Bitcoin rate")
	}

	success := s.Sender.SendRateToUsers(emails, rate)
	if !success {
		return fmt.Errorf("failed to send %d emails", len(emails))
	}

	return nil
}

type EmailHandler struct {
	EmailService *EmailService
}

func NewEmailHandler(storage store.EmailStorage, rateProvider BitcoinRateProvider,emailSender EmailSender) (*EmailHandler, error) {
	emailService := &EmailService{
		Storage: storage,
		Sender:  emailSender,
		RateProvider: rateProvider,
	}

	handler := &EmailHandler{
		EmailService: emailService,
	}

	return handler, nil
}

func (h *EmailHandler) HandleSendEmails(w http.ResponseWriter, r *http.Request) {

	err := h.EmailService.SendEmails()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
