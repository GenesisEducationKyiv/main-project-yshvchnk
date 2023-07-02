package store

import (
	"bitcoin-app/utils"
	"os"
)

type EmailStorage struct {
	StoragePath string
}

func NewEmailStorage(storagePath string) (*EmailStorage, error) {
	es := &EmailStorage{
		StoragePath: storagePath,
	}

	err := es.loadEnv()
	if err != nil {
		return nil, err
	}

	return es, nil
}

func (es *EmailStorage) loadEnv() error {
	utils.LoadEnv()

	es.StoragePath = os.Getenv("STORAGE")
	return nil
}