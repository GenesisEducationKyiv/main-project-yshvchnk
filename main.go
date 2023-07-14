package main

import (
	"bitcoin-app/env"
	"bitcoin-app/handler"
	"bitcoin-app/store"
	rate "bitcoin-app/service/rate"
	send "bitcoin-app/service/send"
	subscribe "bitcoin-app/service/subscribe"
	"bitcoin-app/logger"
	"bitcoin-app/consumer"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
  "github.com/sirupsen/logrus"

)

const storagePath = "../emails.json"

func main() {
	env.LoadEnv()

	port := os.Getenv("PORT")

	router := chi.NewRouter()

	currencyAPI := rate.NewCurrencyAPIProvider()

	currencyRateHandler := handler.NewCurrencyRateHandler(currencyAPI)

	emailStorage, err := store.NewEmailStorage(storagePath)
	if err != nil {
		log.Fatal(err)
	}

	emailSender := &send.EmailSenderPath{
		StoragePath: storagePath,
	}

	emailHandler, err := handler.NewEmailSendHandler(*emailStorage, emailSender)
	if err != nil {
		log.Fatal(err)
	}

	emailServiceSubscribe := subscribe.NewEmailServiceSubscribe()

	router.Get("/api/rate", currencyRateHandler.HandleCurrencyRate)
	router.Post("/api/subscribe", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleSubscribeEmails(w, r, emailServiceSubscribe)
	})
	router.Post("/api/sendEmails", emailHandler.HandleSendEmails)

	log.Println("Server started on port", port)

	serverErr := http.ListenAndServe(":"+port, router)
	if serverErr != nil {
		log.Fatal(serverErr)
	}

	newLogger := logrus.New()
	newLogger.SetLevel(logrus.DebugLevel)

	rabbitmqHook := logger.CreateRabbitMQHook("amqp://guest:guest@localhost:5672/", "logs", "", logrus.ErrorLevel)

  newLogger.AddHook(rabbitmqHook)

	newLogger.Error("error log")
  newLogger.Info("info log")
  newLogger.Debug("debug log")

	consumer.StartConsumer()
}