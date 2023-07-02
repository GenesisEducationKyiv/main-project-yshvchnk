package store

import (
	"encoding/json"
	"os"
)

func (es *EmailStorage) GetEmailsFromFile() ([]string, error) {
	data, err := os.ReadFile(es.StoragePath)
	if err != nil {
		return nil, err
	}

	var emails []string
	err = json.Unmarshal(data, &emails)
	if err != nil {
		return nil, err
	}

	return emails, nil
}