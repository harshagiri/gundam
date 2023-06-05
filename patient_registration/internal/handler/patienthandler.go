package handler

import (
	"encoding/json"
	"net/http"

	"github.com/harshagiri/gundam/patient_registration/internal/model"
	"github.com/harshagiri/gundam/patient_registration/internal/service"
)

type PatientHandler struct {
	patientService *service.PatientService
	// ...
}

func NewPatientHandler(patientService *service.PatientService) *PatientHandler {
	return &PatientHandler{
		patientService: patientService,
		// Initialize other dependencies and configurations for the PatientHandler if needed
	}
}

func (h *PatientHandler) GetPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := h.patientService.GetPatients()
	_ = patients
	if err != nil {
		// Handle the error appropriately
		http.Error(w, "Failed to get patients", http.StatusInternalServerError)
		return
	}

	// Write the response
	// ...
}

func (h *PatientHandler) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the request parameters or URL path
	// ...

	// Call the service method to get the patient by ID
	patient, err := h.patientService.GetPatientByID(id)
	_ = patient
	if err != nil {
		// Handle the error appropriately
		http.Error(w, "Failed to get patient", http.StatusInternalServerError)
		return
	}

	// Write the response
	// ...
}

func (h *PatientHandler) CreatePatient(w http.ResponseWriter, r *http.Request) {
	// Parse the request body and decode it into a Patient struct
	var patient model.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the service method to create the patient
	createdPatient, err := h.patientService.CreatePatient(patient)
	if err != nil {
		http.Error(w, "Failed to create patient", http.StatusInternalServerError)
		return
	}

	// Write the response
	response, err := json.Marshal(createdPatient)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(response)
}

func (h *PatientHandler) GetPatient(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the request parameters or URL path
	// ...

	// Call the service method to get the patient by ID
	patient, err := h.patientService.GetPatientByID(id)
	if err != nil {
		http.Error(w, "Failed to get patient", http.StatusInternalServerError)
		return
	}

	// Write the response
	response, err := json.Marshal(patient)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h *PatientHandler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the request parameters or URL path
	// ...

	// Parse the request body and decode it into a Patient struct
	var patient model.Patient
	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the service method to update the patient
	updatedPatient, err := h.patientService.UpdatePatient(id, patient)
	if err != nil {
		http.Error(w, "Failed to update patient", http.StatusInternalServerError)
		return
	}

	// Write the response
	response, err := json.Marshal(updatedPatient)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h *PatientHandler) DeletePatient(w http.ResponseWriter, r *http.Request) {
	// Parse the ID from the request parameters or URL path
	// ...

	// Call the service method to delete the patient
	err := h.patientService.DeletePatient(id)
	if err != nil {
		http.Error(w, "Failed to delete patient", http.StatusInternalServerError)
		return
	}

	// Write the response
}
