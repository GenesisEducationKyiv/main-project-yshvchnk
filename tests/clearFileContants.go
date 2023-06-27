package test

import(
	"encoding/json"
	"os"
)

func ClearFileContents(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	emails := []string{}
	newContent, err := json.Marshal(emails)
	if err != nil {
		return err
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = file.Write(newContent)
	if err != nil {
		return err
	}

	return nil
}