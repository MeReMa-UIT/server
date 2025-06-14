package errs

import "errors"

var ErrInvalidMedicalRecordStructure = errors.New("Invalid medical record structure (missing required fields or incorrect format and data types)")
var ErrPrimaryDiagnosisMissing = errors.New("Primary diagnosis is missing")
var ErrInvalidAttachmentPrefix = errors.New("Invalid attachment prefix (must start with xray_, ct_, ultrasound_, test_, or other_)")
var ErrRecordNotFound = errors.New("Medical record not found")
