package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/harshagiri/gundam/patient_registration/internal/config"
	"github.com/harshagiri/gundam/patient_registration/internal/handler"
	"github.com/harshagiri/gundam/patient_registration/internal/repository"
	"github.com/harshagiri/gundam/patient_registration/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create the database connection
	db, err := sql.Open("mysql", cfg.DatabaseConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create repositories
	patientRepo := repository.NewPatientRepository(db)
	registrationRepo := repository.NewPatientRegistrationRepository(db)

	// Create services
	patientService := service.NewPatientService(patientRepo)
	registrationService := service.NewPatientRegistrationService(registrationRepo)

	// Create request handlers
	patientHandler := handler.NewPatientHandler(patientService)
	registrationHandler := handler.NewPatientRegistrationHandler(registrationService)

	// Create the router
	router := mux.NewRouter()

	// Register patient-related routes
	router.HandleFunc("/patients", patientHandler.GetPatients).Methods(http.MethodGet)
	router.HandleFunc("/patients", patientHandler.CreatePatient).Methods(http.MethodPost)
	router.HandleFunc("/patients/{id}", patientHandler.GetPatient).Methods(http.MethodGet)
	router.HandleFunc("/patients/{id}", patientHandler.UpdatePatient).Methods(http.MethodPut)
	router.HandleFunc("/patients/{id}", patientHandler.DeletePatient).Methods(http.MethodDelete)

	// Register patient registration-related routes
	router.HandleFunc("/registrations", registrationHandler.GetRegistrations).Methods(http.MethodGet)
	router.HandleFunc("/registrations", registrationHandler.CreateRegistration).Methods(http.MethodPost)
	router.HandleFunc("/registrations/{id}", registrationHandler.GetRegistration).Methods(http.MethodGet)
	router.HandleFunc("/registrations/{id}", registrationHandler.UpdateRegistration).Methods(http.MethodPut)
	router.HandleFunc("/registrations/{id}", registrationHandler.DeleteRegistration).Methods(http.MethodDelete)

	// Start the server
	addr := fmt.Sprintf(":%d", cfg.GetServerPort)
	log.Printf("Server started on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
