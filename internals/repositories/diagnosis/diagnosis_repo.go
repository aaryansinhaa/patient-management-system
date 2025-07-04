package diagnosis_repo

import (
	"database/sql"
	"fmt"

	"github.com/aaryansinhaa/patient-management-system/internals/model"
)

// Package diagnosis_repo provides the implementation of the DiagnosisRepository interface

type DiagnosisStorage struct {
	connection *sql.DB
}

func NewDiagnosisStorage(db *sql.DB) *DiagnosisStorage {
	return &DiagnosisStorage{
		connection: db,
	}
}

func (s *DiagnosisStorage) CreateDiagnosis(diagnosis model.Diagnosis) error {
	query := `INSERT INTO diagnoses (id, patient_id, description, created_at) VALUES ($1, $2, $3, NOW())`
	_, err := s.connection.Exec(query, diagnosis.ID, diagnosis.PatientID, diagnosis.Description)
	if err != nil {
		return fmt.Errorf("failed to create diagnosis: %w", err)
	}
	return nil
}

func (s *DiagnosisStorage) DeleteDiagnosis(id string) (*model.Diagnosis, error) {
	query := `DELETE FROM diagnoses WHERE id = $1 RETURNING id, patient_id,doctor_id, description`
	row := s.connection.QueryRow(query, id)

	var diagnosis model.Diagnosis
	err := row.Scan(&diagnosis.ID, &diagnosis.PatientID, &diagnosis.DoctorID, &diagnosis.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no diagnosis found with id: %s", id)
		}
		return nil, fmt.Errorf("failed to delete diagnosis: %w", err)
	}
	return &diagnosis, nil
}

func (s *DiagnosisStorage) UpdateDiagnosis(diagnosis model.Diagnosis) (*model.Diagnosis, error) {
	query := `UPDATE diagnoses SET patient_id = $1, doctor_id = $2, description = $3, updated_at = NOW()
	          WHERE id = $4 RETURNING id, patient_id, doctor_id, description`
	row := s.connection.QueryRow(query, diagnosis.PatientID, diagnosis.DoctorID, diagnosis.Description, diagnosis.ID)

	var updatedDiagnosis model.Diagnosis
	err := row.Scan(&updatedDiagnosis.ID, &updatedDiagnosis.PatientID, &updatedDiagnosis.DoctorID, &updatedDiagnosis.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to update diagnosis: %w", err)
	}
	return &updatedDiagnosis, nil
}

func (s *DiagnosisStorage) GetDiagnosisByPatientID(patientID string) ([]model.Diagnosis, error) {
	query := `SELECT id, patient_id, doctor_id, description FROM diagnoses WHERE patient_id = $1`
	rows, err := s.connection.Query(query, patientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get diagnoses by patient ID: %w", err)
	}
	defer rows.Close()

	var diagnoses []model.Diagnosis
	for rows.Next() {
		var diagnosis model.Diagnosis
		if err := rows.Scan(&diagnosis.ID, &diagnosis.PatientID, &diagnosis.DoctorID, &diagnosis.Description); err != nil {
			return nil, fmt.Errorf("failed to scan diagnosis: %w", err)
		}
		diagnoses = append(diagnoses, diagnosis)
	}
	return diagnoses, nil
}

