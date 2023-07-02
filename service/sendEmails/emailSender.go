package service

import (
	"bitcoin-app/utils"
	currencyRateGet "bitcoin-app/service/getCurrencyRate"
)

type EmailSenderPath struct {
	StoragePath string
}

func (s *EmailSenderPath) SendEmails(emails []string, rate float64) bool {
	emailService := NewEmailSenderDetails(s.StoragePath)
	return emailService.SendEmails(emails, rate)
}

func (s *EmailSenderPath) GetCurrencyRate() (float64, error) {
	utils.LoadEnv()

	currencyAPI := currencyRateGet.NewCurrencyAPIProvider()

	rate, err := currencyAPI.GetCurrencyRate()
	if err != nil {
		return 0.0, err
	}

	return rate, nil
}