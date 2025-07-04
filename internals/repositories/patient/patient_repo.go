package patient_repo

// Package patient_repo provides the implementation of the PatientRepository interface

import (
	"database/sql"
	"fmt"

	"github.com/aaryansinhaa/patient-management-system/internals/model"
)

type PatientStorage struct {
	connection *sql.DB
}

func NewPatientStorage(db *sql.DB) *PatientStorage {
	return &PatientStorage{
		connection: db,
	}
}

func (s *PatientStorage) CreatePatient(patient model.Patient) error {
	query := `INSERT INTO patients (id, name, age, gender, phone_number) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.connection.Exec(query, patient.ID, patient.Name, patient.Age, patient.Gender, patient.PhoneNumber)
	if err != nil {
		err = fmt.Errorf("failed to create patient: %w", err)
		return err

	}
	return nil
}

func (s *PatientStorage) DeletePatient(id string) (*model.Patient, error) {
	query := `DELETE FROM patients WHERE id = $1 RETURNING id, name, age, phone_number, gender`
	row := s.connection.QueryRow(query, id)
	var patient model.Patient
	err := row.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.PhoneNumber, &patient.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Patient not found
		}
		err = fmt.Errorf("failed to delete patient: %w", err)
		return nil, err
	}
	return &patient, nil
}

func (s *PatientStorage) UpdatePatient(patient model.Patient) (*model.Patient, error) {
	query := `UPDATE patients SET name = $1, age=$2, phone_number=$3, gender=$4, updated_at = NOW()
	          WHERE id = $5 RETURNING id, name, age, phone_number, gender`
	row := s.connection.QueryRow(query, patient.Name, patient.Age, patient.PhoneNumber, patient.Gender, patient.ID)
	var updatedPatient model.Patient
	err := row.Scan(&updatedPatient.ID, &updatedPatient.Name, &updatedPatient.Age, &updatedPatient.PhoneNumber, &updatedPatient.Gender, &updatedPatient.ID)
	if err != nil {
		err = fmt.Errorf("failed to update patient: %w", err)
		return nil, err
	}
	return &updatedPatient, nil
}

func (s *PatientStorage) GetPatientByID(id string) (*model.Patient, error) {
	query := `SELECT id, name, age, phone_number, gender FROM patients WHERE id = $1`
	row := s.connection.QueryRow(query, id)
	var patient model.Patient
	err := row.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.PhoneNumber, &patient.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Patient not found
		}
		err = fmt.Errorf("failed to get patient by ID: %w", err)
		return nil, err
	}
	return &patient, nil
}

func (s *PatientStorage) GetPatientByPhoneNumber(phoneNumber string) (*model.Patient, error) {
	query := `SELECT id, name, age, phone_number, gender FROM patients WHERE phone_number = $1`
	row := s.connection.QueryRow(query, phoneNumber)
	var patient model.Patient
	err := row.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.PhoneNumber, &patient.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Patient not found
		}
		err = fmt.Errorf("failed to get patient by phone number: %w", err)
		return nil, err
	}
	return &patient, nil
}

func (s *PatientStorage) GetAllPatients() ([]model.Patient, error) {
	query := `SELECT id, name, age, phone_number, gender FROM patients`
	rows, err := s.connection.Query(query)
	if err != nil {
		err = fmt.Errorf("failed to get all patients: %w", err)
		return nil, err
	}
	defer rows.Close()
	var patients []model.Patient
	for rows.Next() {
		var patient model.Patient
		err := rows.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.PhoneNumber, &patient.Gender)
		if err != nil {
			err = fmt.Errorf("failed to scan patient row: %w", err)
			return nil, err
		}
		patients = append(patients, patient)
	}
	if err = rows.Err(); err != nil {
		err = fmt.Errorf("error occurred while iterating over patient rows: %w", err)
		return nil, err
	}
	return patients, nil
}

func (s *PatientStorage) GetPatientsByName(name string) ([]model.Patient, error) {
	query := `SELECT id, name, age, phone_number, gender FROM patients WHERE name ILIKE $1`
	rows, err := s.connection.Query(query, "%"+name+"%")
	if err != nil {
		err = fmt.Errorf("failed to get patients by name: %w", err)
		return nil, err
	}
	defer rows.Close()
	var patients []model.Patient
	for rows.Next() {
		var patient model.Patient
		err := rows.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.PhoneNumber, &patient.Gender)
		if err != nil {
			err = fmt.Errorf("failed to scan patient row: %w", err)
			return nil, err
		}
		patients = append(patients, patient)
	}
	if err = rows.Err(); err != nil {
		err = fmt.Errorf("error occurred while iterating over patient rows: %w", err)
		return nil, err
	}
	return patients, nil
}
