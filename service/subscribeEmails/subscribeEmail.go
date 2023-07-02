package service

import (
	"fmt"
)

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