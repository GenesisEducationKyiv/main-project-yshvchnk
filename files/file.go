package files

import (
	"encoding/json"
	"os"
)

//func for checking subscription
func IsEmailSubscribed(email string) bool {
	// read emails from file
	emails, err := GetEmailsFromFile()
	if err != nil {
		return false
	}

	//check if email exists in emails
	for _, e := range emails {
		if e == email {
			return true
		}
	}

	return false
}

//func for saving emails
func SaveEmailToFile(email string) error {
	// read emails from file
	emails, err := GetEmailsFromFile()
	if err != nil {
		emails = []string{}
	}

	//add email and serialize into json
	emails = append(emails, email)
	data, err := json.Marshal(emails)
	if err != nil {
		return err
	}

	//write data into file
	err = os.WriteFile("emails.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetEmailsFromFile() ([]string, error) {
	//read emails from file
	data, err := os.ReadFile("emails.json")
	if err != nil {
		return nil, err
	}

	//deserialize data from file into slice
	var emails []string
	err = json.Unmarshal(data, &emails)
	if err != nil {
		return nil, err
	}

	return emails, nil
}