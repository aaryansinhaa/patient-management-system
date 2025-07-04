package model

import "github.com/google/uuid"

type Diagnosis struct {
	ID          int
	PatientID   uuid.UUID
	DoctorID    uuid.UUID
	Description string
}
