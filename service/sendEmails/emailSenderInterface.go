package service

type EmailSenderInterface interface {
	SendEmail(email string, rate float64) bool
	GetBitcoinRate() (float64, error)
}

