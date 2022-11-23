package uniqueid

import "github.com/google/uuid"

func GenerateUUID() string {
	return uuid.New().String()
}

func NewUUID() string {
	return uuid.New().String()
}
