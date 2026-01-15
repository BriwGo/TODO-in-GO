package app

import "errors"

var EmptyInput error = errors.New("Empty Input")

var WrongArgs error = errors.New("To Many words")
