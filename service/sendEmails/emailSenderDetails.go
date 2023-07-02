package service

type EmailSenderDetails struct {
	StoragePath string
}

func NewEmailSenderDetails(storagePath string) *EmailSenderDetails {
	return &EmailSenderDetails{
		StoragePath: storagePath,
	}
}