package errs

import "errors"

var ErrInvalidExaminationType = errors.New("Invalid choice of examination type")
var ErrInvalidScheduleStatus = errors.New("Invalid choice of schedule status")
