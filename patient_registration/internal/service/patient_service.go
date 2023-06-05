package service

import (
	"github.com/harshagiri/gundam/patient_registration/internal/model"
	"github.com/harshagiri/gundam/patient_registration/internal/repository"
)

type PatientService struct {
	patientRepo repository.PatientRepository
}

func NewPatientService(patientRepo repository.PatientRepository) *PatientService {
	return &PatientService{
		patientRepo: patientRepo,
	}
}

func (s *PatientService) GetPatients() ([]*model.Patient, error) {
	return s.patientRepo.GetAllPatients()
}

func (s *PatientService) GetPatientByID(id int) (*model.Patient, error) {
	return s.patientRepo.GetPatientByID(id)
}

func (s *PatientService) CreatePatient(patient *model.Patient) error {
	// Perform any necessary validations or business logic checks

	// Call the repository to create the patient
	result, err := s.patientRepo.CreatePatient(patient)
	if err != nil {
		// Handle the error
		return err
	}

	return result, nil
}

func (s *PatientService) UpdatePatient(patient *model.Patient) error {
	// Perform any necessary validations or business logic checks

	// Call the repository to update the patient
	err := s.patientRepo.UpdatePatient(patient)
	if err != nil {
		// Handle the error
		return err
	}

	return nil
}

func (s *PatientService) DeletePatient(id int) error {
	// Perform any necessary validations or business logic checks

	// Call the repository to delete the patient
	err := s.patientRepo.DeletePatient(id)
	if err != nil {
		// Handle the error
		return err
	}

	return nil
}
