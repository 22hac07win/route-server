package repository

import "errors"

var ErrNotExists = errors.New("not exist")

var ErrInvalidID = errors.New("invalid id")

var ErrUnAuthorized = errors.New("unauthorized")
