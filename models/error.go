package models

import "errors"

var ErrUsernameNotExists = errors.New("username does not exist")
var ErrPasswordIncorrect = errors.New("password is incorrect")
