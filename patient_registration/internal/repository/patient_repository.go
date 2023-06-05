package repository

import (
	"database/sql"
	"log"

	"github.com/harshagiri/gundam/patient_registration/internal/model"
)

type PatientRepository interface {
	GetAllPatients() ([]*model.Patient, error)
	GetPatientByID(id int) (*model.Patient, error)
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

func (r *patientRepository) CreatePatient(patient *model.Patient) (*model.Patient, error) {
	// Perform any necessary validations or data manipulation

	// Execute the SQL query to insert the patient into the database
	result, err := r.db.Exec("INSERT INTO patients (name, age) VALUES (?, ?)", patient.Name, patient.Age)
	if err != nil {
		// Handle the error
		return nil, err
	}

	// Get the ID of the created patient
	id, err := result.LastInsertId()
	if err != nil {
		// Handle the error
		return nil, err
	}

	// Set the ID of the patient object
	patient.ID = int(id)

	return patient, nil
}

func (r *patientRepository) UpdatePatient(patient *model.Patient) error {
	// Perform any necessary validations or data manipulation

	// Execute the SQL query to update the patient in the database
	_, err := r.db.Exec("UPDATE patients SET name = ?, age = ? WHERE id = ?", patient.Name, patient.Age, patient.ID)
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
