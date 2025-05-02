package models

import "errors"

var ErrCitizenIDNotExists = errors.New("Citizen ID does not exist")
var ErrPasswordIncorrect = errors.New("Password is incorrect")
var ErrCitizenIDExists = errors.New("Citizen ID already exists")
