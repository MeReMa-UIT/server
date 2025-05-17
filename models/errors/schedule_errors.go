package errs

import "errors"

var ErrInvalidExaminationType = errors.New("Invalid choice of examination type, please choose between 1 and 2")
