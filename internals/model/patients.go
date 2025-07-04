package model

import "github.com/google/uuid"

type Patient struct {
	ID          uuid.UUID
	Name        string
	Age         int
	Gender      string
	PhoneNumber string
}
