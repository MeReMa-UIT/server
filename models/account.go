package models

import "time"

type Account struct {
	AccID        uint      `json:"acc_id" db:"acc_id"`
	CitizenID    string    `json:"citizen_id" db:"citizen_id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Phone        string    `json:"phone" db:"phone"`
	Email        *string   `json:"email,omitempty" db:"email"`
	Role         string    `json:"role" db:"role"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	// ExpiresAt time.Time `json:"expires_at"`
}

type SignUpRequest struct {
	CitizenID string  `json:"citizen_id" db:"citizen_id"`
	Password  string  `json:"password"`
	Phone     string  `json:"phone" db:"phone"`
	Email     *string `json:"email,omitempty" db:"email"`
	Role      string  `json:"role" db:"role"`
}

type SignUpResponse struct {
	AccID int `json:"acc_id" db:"acc_id"`
	// PatientID int `json:"patient_id" db:"patient_id"`
	// StaffID  int `json:"staff_id" db:"staff_id"`
}
