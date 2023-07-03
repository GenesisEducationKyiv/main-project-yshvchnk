package main

import (
	"bitcoin-app/handler"
	"bitcoin-app/service"
	"bitcoin-app/store"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

const storagePath = "../emails.json"

type EmailSender struct {
	StoragePath string
}

func (s *EmailSender) SendRateToUsers(emails []string, rate float64) bool {
	emailService := service.NewEmailSenderDetails(s.StoragePath)
	return emailService.SendEmails(emails, rate)
}

func (s *EmailSender) GetBitcoinRate() (float64, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	coinGeckoAPI := service.NewCoinGeckoAPI()

	rate, err := coinGeckoAPI.GetBitcoinRate()
	if err != nil {
		return 0.0, err
	}

	return rate, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	storage, err := store.NewEmailStorage(storagePath)
	if err != nil {
		log.Fatal("Failed to create email storage:", err)
	}

	emailServiceInst := service.NewEmailService()

	router := chi.NewRouter()

	bitcoinAPI := service.NewCoinGeckoAPI()

	bitcoinRateHandler := handler.NewBitcoinRateHandler(bitcoinAPI)

	emailSender := &EmailSender{
		StoragePath: storagePath,
	}

	emailService := &handler.EmailService{
		Storage:       *storage,
		Sender:        emailSender,
		RateProvider:  bitcoinAPI,
	}

	emailHandler := &handler.EmailHandler{
		EmailService:  emailService,
	}

	router.Get("/api/rate", bitcoinRateHandler.HandleRate)
	router.Post("/api/subscribe", handler.HandleSubscribe(emailServiceInst))
	router.Post("/api/sendEmails", emailHandler.HandleSendEmails)

	log.Println("Server started on port", port)

	serverErr := http.ListenAndServe(":"+port, router)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}