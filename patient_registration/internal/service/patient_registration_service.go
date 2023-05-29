package service

import (
	"../../internal/model"
	"../../internal/repository"
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
