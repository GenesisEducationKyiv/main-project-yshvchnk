package service

import (
	"bitcoin-app/store"
)

const storagePath = "../emails.json"

type EmailStorageInterface interface {
	IsEmailSubscribed(email string) (bool, error)
	SaveEmailToFile(email string) error
}

type EmailServiceSubscribe struct {
	storage EmailStorageInterface
}

func NewEmailServiceSubscribe() *EmailServiceSubscribe {
	storagePath := storagePath
	storage, _ := store.NewEmailStorage(storagePath)
	return &EmailServiceSubscribe{
		storage: storage,
	}
}