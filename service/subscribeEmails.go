package service

import (
	"bitcoin-app/file"
	"errors"
	"fmt"
	"log"
)

var (
	ErrSubscriptionCheckFailed = errors.New("subscription check failed")
	ErrEmailAlreadySubscribed  = errors.New("email already subscribed")
	ErrFailedToSaveEmail       = errors.New("failed to save email")
)

func SubscribeEmail(email string) error {

	es, err := file.NewEmailStorage("../emails.json")
	if err != nil {
			log.Fatal(err)
	}

	subscribed, err := es.IsEmailSubscribed(email)
	if err != nil {
		return fmt.Errorf("failed to check subscription: %s: %s", ErrSubscriptionCheckFailed, err.Error())
	}

	if subscribed {
		return ErrEmailAlreadySubscribed
	}

	err = es.SaveEmailToFile(email)
	if err != nil {
		return fmt.Errorf("failed to save email: %s: %s", ErrFailedToSaveEmail, err.Error())
	}

	return nil
}