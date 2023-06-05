package repository

import (
	"database/sql"
	"log"

	"github.com/harshagiri/gundam/patient_registration/internal/model"
)

type PatientRegistrationRepository interface {
	GetAllRegistrations() ([]*model.PatientRegistration, error)
	GetRegistrationByID(id int) (*model.PatientRegistration, error)
	CreateRegistration(patient_registration *model.PatientRegistration) (int64, error)
	UpdateRegistration(patient_registration *model.PatientRegistration) (int64, error)
	DeleteRegistration(id int) error
}

type patientRegistrationRepository struct {
	db *sql.DB
}

func NewPatientRegistrationRepository(db *sql.DB) PatientRegistrationRepository {
	return &patientRegistrationRepository{
		db: db,
	}
}

func (r *patientRegistrationRepository) GetAllRegistrations() ([]*model.PatientRegistration, error) {
	rows, err := r.db.Query("SELECT * FROM patient_registrations")
	if err != nil {
		log.Printf("Failed to retrieve patient registrations: %v", err)
		return nil, err
	}
	defer rows.Close()

	registrations := []*model.PatientRegistration{}
	for rows.Next() {
		registration := &model.PatientRegistration{}
		err := rows.Scan(
			&registration.ID,
			&registration.RegistrationDate,
			&registration.PatientID,
			&registration.FirstName,
			&registration.LastName,
			&registration.DateOfBirth,
			&registration.Gender,
			&registration.Email,
			&registration.PhoneNumber,
			&registration.Address,
			&registration.CreatedAt,
		)
		if err != nil {
			log.Printf("Failed to scan patient registration row: %v", err)
			return nil, err
		}
		registrations = append(registrations, registration)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error occurred while iterating patient registration rows: %v", err)
		return nil, err
	}

	return registrations, nil
}

func (r *patientRegistrationRepository) GetRegistrationByID(id int) (*model.PatientRegistration, error) {
	query := "SELECT * FROM patient_registrations WHERE id = ?"
	row := r.db.QueryRow(query, id)

	registration := &model.PatientRegistration{}
	err := row.Scan(
		&registration.ID,
		&registration.RegistrationDate,
		&registration.PatientID,
		&registration.FirstName,
		&registration.LastName,
		&registration.DateOfBirth,
		&registration.Gender,
		&registration.Email,
		&registration.PhoneNumber,
		&registration.Address,
		&registration.CreatedAt,
	)
	if err != nil {
		log.Printf("Failed to retrieve patient registration with ID %d: %v", id, err)
		return nil, err
	}

	return registration, nil
}

// Implement other repository methods such as CreateRegistration, UpdateRegistration, DeleteRegistration, etc.
func (r *patientRegistrationRepository) CreateRegistration(patient_registration *model.PatientRegistration) (int64, error) {
	// Execute the SQL query to insert the patient into the database
	result, err := r.db.Exec("INSERT INTO patients (first_name, last_name, date_of_birth, gender, email, phone_number, address, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		patient_registration.FirstName, patient_registration.LastName, patient_registration.DateOfBirth,
		patient_registration.Gender, patient_registration.Email, patient_registration.PhoneNumber,
		patient_registration.Address, patient_registration.CreatedAt)
	if err != nil {
		log.Printf("Failed to insert patient: %v", err)
		return 0, err
	}

	// Get the ID of the inserted patientss
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Failed to get last inserted ID: %v", err)
		return 0, err
	}

	return id, nil
}

func (r *patientRegistrationRepository) UpdateRegistration(patient_registration *model.PatientRegistration) (int64, error) {
	// Perform any necessary validations or data manipulation

	// Execute the SQL query to update the patient in the database
	_, err := r.db.Exec("UPDATE patients SET first_name = ?, last_name = ?, date_of_birth = ?, gender = ?, email = ?, phone_number = ?, address = ?, created_at = ? WHERE id = ?",
		patient_registration.FirstName, patient_registration.LastName, patient_registration.DateOfBirth,
		patient_registration.Gender, patient_registration.Email, patient_registration.PhoneNumber,
		patient_registration.Address, patient_registration.CreatedAt, patient_registration.ID)
	if err != nil {
		// Handle the error
		return 0, err
	}

	return 0, nil
}

func (r *patientRegistrationRepository) DeleteRegistration(id int) error {
	// Perform any necessary validations or data manipulation

	// Execute the SQL query to delete the patient from the database
	_, err := r.db.Exec("DELETE FROM patients WHERE id = ?", id)
	if err != nil {
		// Handle the error
		return err
	}

	return nil
}
