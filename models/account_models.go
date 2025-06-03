package models

import "time"

type AccountInfo struct {
	AccID     int       `json:"acc_id" db:"acc_id"`
	CitizenID string    `json:"citizen_id" db:"citizen_id"`
	Phone     string    `json:"phone" db:"phone"`
	Email     string    `json:"email" db:"email"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type LoginRequest struct {
	Identifier string `json:"id" db:"id"`
	Password   string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type InitRegistrationRequest struct {
	CitizenID string `json:"citizen_id" db:"citizen_id"`
}

type AccountRegistrationResponse struct {
	// JWT token. If acc ID = -1, token will allow user to register new account, otherwise token will allow user to add new patient or staff
	Token string `json:"token"`

	// Account ID (-1 means account is not registered yet)
	AccID int `json:"acc_id"`
}

type AccountRegistrationRequest struct {
	CitizenID string `json:"citizen_id" db:"citizen_id"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" db:"email"`
	Role      string `json:"role" db:"role"`
}

type PatientRegistrationRequest struct {
	AccID                      int       `json:"acc_id" db:"acc_id"`
	FullName                   string    `json:"full_name" db:"full_name"`
	DateOfBirth                time.Time `json:"date_of_birth" db:"date_of_birth"`
	Gender                     string    `json:"gender" db:"gender"`
	Ethnicity                  string    `json:"ethnicity" db:"ethnicity"`
	Nationality                string    `json:"nationality" db:"nationality"`
	Address                    string    `json:"address" db:"address"`
	HealthInsuranceExpiredDate time.Time `json:"health_insurance_expired_date" db:"health_insurance_expired_date"`
	HealthInsuranceNumber      string    `json:"health_insurance_number" db:"health_insurance_number"`
	EmergencyContactInfo       string    `json:"emergency_contact_info" db:"emergency_contact_info"`
}

type StaffRegistrationRequest struct {
	AccID       int       `json:"acc_id" db:"acc_id"`
	FullName    string    `json:"full_name" db:"full_name"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
	Gender      string    `json:"gender" db:"gender"`
	Department  string    `json:"department" db:"department"`
}

type AccountRecoverRequest struct {
	CitizenID string `json:"citizen_id" db:"citizen_id"`
	Email     string `json:"email" db:"email"`
}

type AccountRecoverConfirmRequest struct {
	CitizenID string `json:"citizen_id" db:"citizen_id"`
	OTP       string `json:"otp" db:"otp"`
}

type AccountRecoverConfirmResponse struct {
	Token string `json:"token"`
}

type PasswordResetRequest struct {
	NewPassword string `json:"new_password"`
}

type UpdateAccountInfoRequest struct {
	// "password", "email", "phone", "citizen_id" are possible choices
	Field    string `json:"field"`
	NewValue string `json:"new_value"`
	Password string `json:"password"`
}
