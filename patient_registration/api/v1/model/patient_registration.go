package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../internal/service"
)

type PatientRegistrationAPI struct {
	registrationService *service.PatientRegistrationService
}

func NewPatientRegistrationAPI(registrationService *service.PatientRegistrationService) *PatientRegistrationAPI {
	return &PatientRegistrationAPI{
		registrationService: registrationService,
	}
}

func (api *PatientRegistrationAPI) GetAllRegistrations(w http.ResponseWriter, r *http.Request) {
	registrations, err := api.registrationService.GetRegistrations()
	if err != nil {
		log.Printf("Failed to get patient registrations: %v", err)
		http.Error(w, "Failed to get patient registrations", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, registrations)
}

func (api *PatientRegistrationAPI) GetRegistrationByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid registration ID", http.StatusBadRequest)
		return
	}

	registration, err := api.registrationService.GetRegistrationByID(id)
	if err != nil {
		log.Printf("Failed to get patient registration with ID %d: %v", id, err)
		http.Error(w, "Failed to get patient registration", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, registration)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
