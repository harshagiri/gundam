package repository

import (
	"database/sql"
	"log"

	"github.com/harshagiri/gundam/patient_registration/internal/model"
)

type PatientRepository interface {
	// Define the methods of the PatientRepository interface
}

type patientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) PatientRepository {
	return &patientRepository{
		db: db,
	}
}

func (r *patientRepository) GetAllPatients() ([]*model.Patient, error) {
	rows, err := r.db.Query("SELECT * FROM patients")
	if err != nil {
		log.Printf("Failed to retrieve patients: %v", err)
		return nil, err
	}
	defer rows.Close()

	patients := []*model.Patient{}
	for rows.Next() {
		patient := &model.Patient{}
		err := rows.Scan(
			&patient.ID,
			&patient.FirstName,
			&patient.LastName,
			&patient.DateOfBirth,
			&patient.Gender,
			&patient.Email,
			&patient.PhoneNumber,
			&patient.Address,
			&patient.CreatedAt,
		)
		if err != nil {
			log.Printf("Failed to scan patient row: %v", err)
			return nil, err
		}
		patients = append(patients, patient)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error occurred while iterating patient rows: %v", err)
		return nil, err
	}

	return patients, nil
}

func (r *patientRepository) GetPatientByID(id int) (*model.Patient, error) {
	query := "SELECT * FROM patients WHERE id = ?"
	row := r.db.QueryRow(query, id)

	patient := &model.Patient{}
	err := row.Scan(
		&patient.ID,
		&patient.FirstName,
		&patient.LastName,
		&patient.DateOfBirth,
		&patient.Gender,
		&patient.Email,
		&patient.PhoneNumber,
		&patient.Address,
		&patient.CreatedAt,
	)
	if err != nil {
		log.Printf("Failed to retrieve patient with ID %d: %v", id, err)
		return nil, err
	}

	return patient, nil
}

// Implement other repository methods such as CreatePatient, UpdatePatient, DeletePatient, etc.
