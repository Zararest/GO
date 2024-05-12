package errors

import "errors"

var (
	// ErrFileNotExists using when file not exists for sfss service.
	ErrFileNotExists = errors.New("sfss file not exists")
	// ErrFileAlreadyExists using when file already exists in system.
	ErrFileAlreadyExists = errors.New("sfss file alredy exists")
	// ErrInternalError using if something went wrong.
	ErrInternalError = errors.New("sfss internal error")
)
