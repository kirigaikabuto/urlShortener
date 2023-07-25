package storage

import "errors"

var (
	ErrURLNotFound = errors.New("ulr not found")
	ErrURLExists   = errors.New("url exists")
)
