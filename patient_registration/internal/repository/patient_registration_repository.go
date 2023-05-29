package repository

import (
	"database/sql"
	"log"

	"../../internal/model"
)

type PatientRegistrationRepository struct {
	db *sql.DB
}

func NewPatientRegistrationRepository(db *sql.DB) *PatientRegistrationRepository {
	return &PatientRegistrationRepository{
		db: db,
	}
}

func (r *PatientRegistrationRepository) GetAllRegistrations() ([]*model.PatientRegistration, error) {
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

func (r *PatientRegistrationRepository) GetRegistrationByID(id int) (*model.PatientRegistration, error) {
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
