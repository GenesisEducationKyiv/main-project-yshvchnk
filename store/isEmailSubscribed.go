package store

import (
	"github.com/pkg/errors"
)

func (es *EmailStorage) IsEmailSubscribed(email string) (bool, error) {
	emails, err := es.GetEmailsFromFile()
	if err != nil {
		return false, errors.Wrap(err, "Failed to load email addresses")
	}

	for _, e := range emails {
		if e == email {
			return true, nil
		}
	}

	return false, nil
}