package service

import (
	"bitcoin-app/store"
	"fmt"
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

func (es *EmailServiceSubscribe) SubscribeEmail(email string) error {
	subscribed, err := es.storage.IsEmailSubscribed(email)
	if err != nil {
		return fmt.Errorf("failed to check subscription: %s: %s", ErrSubscriptionCheckFailed, err.Error())
	}

	if subscribed {
		return ErrEmailAlreadySubscribed
	}

	err = es.storage.SaveEmailToFile(email)
	if err != nil {
		return fmt.Errorf("failed to save email: %s: %s", ErrFailedToSaveEmail, err.Error())
	}

	return nil
}