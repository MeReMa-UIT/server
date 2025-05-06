package errors

import (
	"errors"
)

var ErrAccountNotExist = errors.New("Account does not exist")
var ErrPasswordIncorrect = errors.New("Password is incorrect")
var ErrCitizenIDExists = errors.New("Citizen ID already exists")
var ErrEmailOrPhoneAlreadyUsed = errors.New("Email or phone number is already used")

var ErrWrongOTP = errors.New("Wrong OTP")
var ErrExpiredOTP = errors.New("Expired OTP")

var ErrInvalidToken = errors.New("Invalid token")
var ErrExpiredToken = errors.New("Expired token")
var ErrMalformedToken = errors.New("Malformed token")
var ErrPermissionDenied = errors.New("Permission denied. You are not allowed to perform this action")
