package repository

import (
	"database/sql"
	"log"

	"github.com/harshagiri/gundam/patient_registration/internal/model"
)

type PatientRepository interface {
	GetAllPatients() ([]*model.Patient, error)
	GetPatientByID(id int) (*model.Patient, error)
	CreatePatient(patient *model.Patient) (int64, error)
	UpdatePatient(patient *model.Patient) (int64, error)
	DeletePatient(id int) error
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

func (r *patientRepository) CreatePatient(patient *model.Patient) (int64, error) {
	// Execute the SQL query to insert the patient into the database
	result, err := r.db.Exec("INSERT INTO patients (first_name, last_name, date_of_birth, gender, email, phone_number, address, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		patient.FirstName, patient.LastName, patient.DateOfBirth,
		patient.Gender, patient.Email, patient.PhoneNumber,
		patient.Address, patient.CreatedAt)
	if err != nil {
		log.Printf("Failed to insert patient: %v", err)
		return 0, err
	}

	// Get the ID of the inserted patient
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Failed to get last inserted ID: %v", err)
		return 0, err
	}

	return id, nil
}

func (r *patientRepository) UpdatePatient(patient *model.Patient) error {
	// Perform any necessary validations or data manipulation

	// Execute the SQL query to update the patient in the database
	_, err := r.db.Exec("UPDATE patients SET first_name = ?, last_name = ?, date_of_birth = ?, gender = ?, email = ?, phone_number = ?, address = ?, created_at = ? WHERE id = ?",
		patient.FirstName, patient.LastName, patient.DateOfBirth,
		patient.Gender, patient.Email, patient.PhoneNumber,
		patient.Address, patient.CreatedAt, patient.ID)
	if err != nil {
		// Handle the error
		return err
	}

	return nil
}

func (r *patientRepository) DeletePatient(id int) error {
	// Perform any necessary validations or data manipulation

	// Execute the SQL query to delete the patient from the database
	_, err := r.db.Exec("DELETE FROM patients WHERE id = ?", id)
	if err != nil {
		// Handle the error
		return err
	}

	return nil
}
