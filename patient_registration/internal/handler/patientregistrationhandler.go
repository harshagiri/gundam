package handler

import (
	"encoding/json"
	"net/http"

	"github.com/harshagiri/gundam/patient_registration/internal/model"
	"github.com/harshagiri/gundam/patient_registration/internal/service"
)

type PatientRegistrationHandler struct {
	registrationService *service.PatientRegistrationService
	// ...
}

func NewPatientRegistrationHandler(registrationService *service.PatientRegistrationService) *PatientRegistrationHandler {
	return &PatientRegistrationHandler{
		registrationService: registrationService,
		// Initialize other dependencies and configurations for the PatientRegistrationHandler if needed
	}
}

func (h *PatientRegistrationHandler) GetRegistrations(w http.ResponseWriter, r *http.Request) {
	registrations, err := h.registrationService.GetRegistrations()
	_ = registrations
	if err != nil {
		http.Error(w, "Failed to get registrations", http.StatusInternalServerError)
		return
	}

	// Write the response
	// ...
}

func (h *PatientRegistrationHandler) GetRegistrationByID(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the request parameters or URL path
	// ...

	// Call the service method to get the registration by ID
	registration, err := h.registrationService.GetRegistrationByID(id)
	_ = registration
	if err != nil {
		http.Error(w, "Failed to get registration", http.StatusInternalServerError)
		return
	}

	// Write the response
	// ...
}

func (h *PatientRegistrationHandler) CreateRegistration(w http.ResponseWriter, r *http.Request) {
	// Parse the request body and decode it into a PatientRegistration struct
	var registration model.PatientRegistration
	err := json.NewDecoder(r.Body).Decode(&registration)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the service method to create the registration
	createdRegistration, err := h.registrationService.CreateRegistration(registration)
	if err != nil {
		http.Error(w, "Failed to create registration", http.StatusInternalServerError)
		return
	}

	// Write the response
	response, err := json.Marshal(createdRegistration)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(response)
}

func (h *PatientRegistrationHandler) GetRegistration(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the request parameters or URL path
	// ...

	// Call the service method to get the registration by ID
	registration, err := h.registrationService.GetRegistrationByID(id)
	if err != nil {
		http.Error(w, "Failed to get registration", http.StatusInternalServerError)
		return
	}

	// Write the response
	response, err := json.Marshal(registration)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h *PatientRegistrationHandler) UpdateRegistration(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the request parameters or URL path
	// ...

	// Parse the request body and decode it into a PatientRegistration struct
	var registration model.PatientRegistration
	err := json.NewDecoder(r.Body).Decode(&registration)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the service method to update the registration
	updatedRegistration, err := h.registrationService.UpdateRegistration(id, registration)
	if err != nil {
		http.Error(w, "Failed to update registration", http.StatusInternalServerError)
		return
	}

	// Write the response
	response, err := json.Marshal(updatedRegistration)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h *PatientRegistrationHandler) DeleteRegistration(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the request parameters or URL path
	// ...

	// Call the service method to delete the registration
	err := h.registrationService.DeleteRegistration(id)
	if err != nil {
		http.Error(w, "Failed to delete registration", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
}

// Implement other handler methods as needed...
