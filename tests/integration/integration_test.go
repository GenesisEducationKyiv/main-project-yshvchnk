package test

import (
	"bitcoin-app/store"
	"bitcoin-app/service/subscribeEmails"
	"bitcoin-app/tests"
	"testing"
)

func TestServiceAndDatabaseWillReturnSuccess(t *testing.T) {

	err := test.ClearFileContents("emails.json")
	if err != nil {
		t.Fatalf("Error cleaning up emails.json: %s", err)
	}

	testEmail := "test2@example.com"

	es, err := store.NewEmailStorage("emails.json")
	if err != nil {
			t.Fatal(err)
	}

	emailService := service.NewEmailServiceSubscribe()

	err = emailService.SubscribeEmail(testEmail)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	subscribed, err := es.IsEmailSubscribed(testEmail);
	if err != nil {
		t.Errorf("failed to check subscription: %v", err)
	}

	if !subscribed {
		t.Error("email was not added")
	}

}