package model

type PatientRegistration struct {
	ID               int    `json:"id"`
	RegistrationDate string `json:"registration_date"`
	PatientID        int    `json:"patient_id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	DateOfBirth      string `json:"date_of_birth"`
	Gender           string `json:"gender"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	Address          string `json:"address"`
	CreatedAt        string `json:"created_at"`
}
