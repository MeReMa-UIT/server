package models

import "errors"

var ErrAccountNotExist = errors.New("Account does not exist")
var ErrPasswordIncorrect = errors.New("Password is incorrect")
var ErrCitizenIDExists = errors.New("Citizen ID already exists")
var ErrEmailDoesNotMatchCitizenID = errors.New("Email and citizen ID do not match")

var ErrWrongOTP = errors.New("Wrong OTP")
var ErrExpiredOTP = errors.New("Expired OTP")
var ErrUnverifiedOTP = errors.New("OTP not verified")
