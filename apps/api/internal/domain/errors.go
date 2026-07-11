package domain

import "errors"

var (
	ErrNotFound     = errors.New("entity not found")
	ErrInvalidInput = errors.New("invalid input")
)
