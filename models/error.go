package models

import "errors"

var ErrUsernameNotExists = errors.New("Username does not exist")
var ErrPasswordIncorrect = errors.New("Password is incorrect")
var ErrCitizenIDExists = errors.New("Citizen ID already exists")
