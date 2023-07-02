package store

import (
	"encoding/json"
	"os"
)

func (es *EmailStorage) SaveEmailToFile(email string) error {
	emails, err := es.GetEmailsFromFile()
	if err != nil {
		emails = []string{}
	}

	emails = append(emails, email)
	data, err := json.Marshal(emails)
	if err != nil {
		return err
	}

	err = os.WriteFile(es.StoragePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}