package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../internal/service"
)

type PatientAPI struct {
	patientService *service.PatientService
}

func NewPatientAPI(patientService *service.PatientService) *PatientAPI {
	return &PatientAPI{
		patientService: patientService,
	}
}

func (api *PatientAPI) GetAllPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := api.patientService.GetPatients()
	if err != nil {
		log.Printf("Failed to get patients: %v", err)
		http.Error(w, "Failed to get patients", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, patients)
}

func (api *PatientAPI) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := api.patientService.GetPatientByID(id)
	if err != nil {
		log.Printf("Failed to get patient with ID %d: %v", id, err)
		http.Error(w, "Failed to get patient", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, patient)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
