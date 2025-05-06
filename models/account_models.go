package models

import "time"

// type Account struct {
// 	AccID        uint      `json:"acc_id" db:"acc_id"`
// 	CitizenID    string    `json:"citizen_id" db:"citizen_id"`
// 	PasswordHash string    `json:"-" db:"password_hash"`
// 	Phone        string    `json:"phone" db:"phone"`
// 	Email        *string   `json:"email,omitempty" db:"email"`
// 	Role         string    `json:"role" db:"role"`
// 	CreatedAt    time.Time `json:"created_at" db:"created_at"`
// }

type LoginRequest struct {
	Identifier string `json:"id" db:"id"`
	Password   string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	// ExpiresAt time.Time `json:"expires_at"`
}

type AccountRegisterRequest struct {
	CitizenID string  `json:"citizen_id" db:"citizen_id"`
	Password  string  `json:"password"`
	Phone     string  `json:"phone" db:"phone"`
	Email     *string `json:"email,omitempty" db:"email"`
	Role      string  `json:"role" db:"role"`
}

type PatientRegisterRequest struct {
	AccountRegisterRequest
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

type StaffRegisterRequest struct {
	AccountRegisterRequest
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
