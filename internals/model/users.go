package model

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID
	Name        string
	Role        string
	Username    string
	Password    string
	PhoneNumber string
}
