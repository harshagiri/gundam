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

// Implement other service methods such as CreatePatient, UpdatePatient, DeletePatient, etc.
