package main

import (
	rateCurrencyHandler "bitcoin-app/handler/handleCurrencyRate"
	sendEmailsHandler "bitcoin-app/handler/handleSendEmails"
	subscribeEmailsHandler "bitcoin-app/handler/handleSubscribeEmails"
	currencyRateGet "bitcoin-app/service/getCurrencyRate"
	emailSend "bitcoin-app/service/sendEmails"
	"bitcoin-app/utils"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
)

const storagePath = "../emails.json"

func main() {
	utils.LoadEnv()

	port := os.Getenv("PORT")

	router := chi.NewRouter()

	currencyAPI := currencyRateGet.NewCurrencyAPIProvider()

	currencyRateHandler := rateCurrencyHandler.NewCurrencyRateHandler(currencyAPI)

	emailSender := &emailSend.EmailSenderPath{
		StoragePath: storagePath,
	}

	emailHandler, err := sendEmailsHandler.NewEmailSendHandler(storagePath, currencyAPI, emailSender)
	if err != nil {
		log.Fatal(err)
	}

	router.Get("/api/rate", currencyRateHandler.HandleCurrencyRate)
	router.Post("/api/subscribe", subscribeEmailsHandler.HandleSubscribeEmails)
	router.Post("/api/sendEmails", emailHandler.HandleSendEmails)

	log.Println("Server started on port", port)

	serverErr := http.ListenAndServe(":"+port, router)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}