package service

import (
	"errors"
)

var (
	ErrSubscriptionCheckFailed = errors.New("subscription check failed")
	ErrEmailAlreadySubscribed  = errors.New("email already subscribed")
	ErrFailedToSaveEmail       = errors.New("failed to save email")
)