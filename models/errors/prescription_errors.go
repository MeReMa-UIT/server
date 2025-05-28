package errs

import "errors"

var ErrInvalidDosage = errors.New("Invalid dosage values provided")
var ErrWrongDosageCalulation = errors.New("Total dosage does not match the sum of individual dosages multiplied by duration")
