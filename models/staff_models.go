package models

import "time"

type StaffInfo struct {
	StaffID     int       `json:"staff_id" db:"staff_id"`
	FullName    string    `json:"full_name" db:"full_name"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
	Gender      string    `json:"gender" db:"gender"`
	Department  string    `json:"department" db:"department"`
}

type StaffInfoUpdateRequest struct {
	FullName    string    `json:"full_name" db:"full_name"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
	Gender      string    `json:"gender" db:"gender"`
	Department  string    `json:"department" db:"department"`
}
