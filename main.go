package main

import (
	"bitcoin-app/handler"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := chi.NewRouter()

	storagePath := "../emails.json"
	emailHandler, err := handler.NewEmailHandler(storagePath)
	if err != nil {
		log.Fatal(err)
	}

	router.Get("/api/rate", handler.HandleRate)
	router.Post("/api/subscribe", handler.HandleSubscribe)
	router.Post("/api/sendEmails", emailHandler.HandleSendEmails)

	log.Println("Server started on port", port)

	serverErr := http.ListenAndServe(":"+port, router)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}