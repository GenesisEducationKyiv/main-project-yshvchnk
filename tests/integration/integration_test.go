package test

import (
	"bitcoin-app/file"
	"bitcoin-app/service"
	"bitcoin-app/tests"
	"log"
	"testing"
)

func TestServiceAndDatabaseWillReturnSuccess(t *testing.T) {

	err := test.ClearFileContents("emails.json")
	if err != nil {
		t.Fatalf("Error cleaning up emails.json: %s", err)
	}

	testEmail := "test2@example.com"

	if err := service.SubscribeEmail(testEmail); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	es, err := file.NewEmailStorage("emails.json")
	if err != nil {
			log.Fatal(err)
	}

	subscribed, err := es.IsEmailSubscribed(testEmail);
	if err != nil {
		t.Errorf("failed to check subscription: %v", err)
	}

	if !subscribed {
		t.Error("email was not added")
	}

}