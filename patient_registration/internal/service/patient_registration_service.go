package service

import (
	"github.com/harshagiri/gundam/patient_registration/internal/model"
	"github.com/harshagiri/gundam/patient_registration/internal/repository"
)

type PatientRegistrationService struct {
	registrationRepo repository.PatientRegistrationRepository
}

func NewPatientRegistrationService(registrationRepo repository.PatientRegistrationRepository) *PatientRegistrationService {
	return &PatientRegistrationService{
		registrationRepo: registrationRepo,
	}
}

func (s *PatientRegistrationService) GetRegistrations() ([]*model.PatientRegistration, error) {
	return s.registrationRepo.GetAllRegistrations()
}

func (s *PatientRegistrationService) GetRegistrationByID(id int) (*model.PatientRegistration, error) {
	return s.registrationRepo.GetRegistrationByID(id)
}

// Implement other service methods such as CreateRegistration, UpdateRegistration, DeleteRegistration, etc.
func (s *PatientRegistrationService) CreateRegistration(registration *model.PatientRegistration) (int64, error) {
	// Perform any necessary validations or business logic checks

	// Call the repository to create the patient
	result, err := s.registrationRepo.CreateRegistration(registration)
	if err != nil {
		// Handle the error
		return 0, err
	}

	return result, nil
}

func (s *PatientRegistrationService) UpdateRegistration(registration *model.PatientRegistration) (int64, error) {
	// Perform any necessary validations or business logic checks

	// Call the repository to update the patient
	result, err := s.registrationRepo.UpdateRegistration(registration)
	if err != nil {
		// Handle the error
		return 0, err
	}

	return result, nil
}

func (s *PatientRegistrationService) DeleteRegistration(id int) error {
	// Perform any necessary validations or business logic checks

	// Call the repository to delete the patient
	err := s.registrationRepo.DeleteRegistration(id)
	if err != nil {
		// Handle the error
		return err
	}

	return nil
}
