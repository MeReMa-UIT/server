package errs

import "errors"

var ErrInvalidDosage = errors.New("Invalid dosage values or duration days were provided")
var ErrWrongDosageCalulation = errors.New("Total dosage does not match the sum of individual dosages multiplied by duration")
var ErrReceivedPrescription = errors.New("Prescription has already been received")
var ErrPrescriptionNotFound = errors.New("Prescription not found")
var ErrPrescriptionDetailNotFound = errors.New("Prescription detail not found")
