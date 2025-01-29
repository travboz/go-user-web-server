package main

import "errors"

var (
	ErrUserDoesNotExist = errors.New("user does not exist")
	ErrNameRequired     = errors.New("name is required")
)
