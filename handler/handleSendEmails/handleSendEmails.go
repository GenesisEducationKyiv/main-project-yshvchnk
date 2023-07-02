package handler

import (
	"bitcoin-app/store"
	"bitcoin-app/service/sendEmails"
	"net/http"
	"github.com/pkg/errors"
)

type EmailSendHandler struct {
	EmailService *service.EmailSendService
}

func NewEmailSendHandler(storagePath string, rateProvider CurrencyRateProvider,emailSender service.EmailSender) (*EmailSendHandler, error) {
	storage, err := store.NewEmailStorage(storagePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create email storage")
	}

	emailService := &service.EmailSendService{
		Storage: *storage,
		Sender:  emailSender,
	}

	handler := &EmailSendHandler{
		EmailService: emailService,
	}

	return handler, nil
}

func (h *EmailSendHandler) HandleSendEmails(w http.ResponseWriter, r *http.Request) {
	err := h.EmailService.SendEmails()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
