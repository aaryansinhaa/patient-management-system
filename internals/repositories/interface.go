package repositories

import "github.com/aaryansinhaa/patient-management-system/internals/model"

type UserRepository interface {
	CreateUser(user model.User) error
	DeleteUser(id string) (*model.User, error)
	UpdateUser(user model.User) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	GetAllUsersByRole(role string) ([]model.User, error)
	GetUserByPhoneNumber(phoneNumber string) (*model.User, error)
	GetUsernameAndPasswordById(id string) (string, string, error)
}

type PatientRepository interface {
	CreatePatient(patient model.Patient) error
	DeletePatient(id string) (*model.Patient, error)
	UpdatePatient(patient model.Patient) (*model.Patient, error)
	GetPatientByID(id string) (*model.Patient, error)
	GetPatientByPhoneNumber(phoneNumber string) (*model.Patient, error)
	GetAllPatients() ([]model.Patient, error)
	GetAllPatientsByName(name string) ([]model.Patient, error)
}

type DiagnosisRepository interface {
	CreateDiagnosis(diagnosis model.Diagnosis) error
	DeleteDiagnosis(id string) (*model.Diagnosis, error)
	UpdateDiagnosis(diagnosis model.Diagnosis) (*model.Diagnosis, error)
	GetDiagnosisByPatientID(patientID string) ([]model.Diagnosis, error)
}
