package errs

import "errors"

var ErrInvalidCompileType = errors.New("Invalid compile type provided. Valid options are: time, doctor, diagnosis")
var ErrInvalidTimeUnit = errors.New("Invalid time unit provided. Valid options are: day, week, month, year")
var ErrInvalidTimestamp = errors.New("Invalid timestamp provided.")
