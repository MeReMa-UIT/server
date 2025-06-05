package errs

import "errors"

var ErrInvalidMedicalRecordStructure = errors.New("Invalid medical record structure (missing required fields or incorrect format and data types)")
