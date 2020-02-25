package models

import "errors"

var (
	// ErrBadParamInput generic error when entity operation cannot be done due to invalid/missing input. Should be set to identify the input param and validation message
	ErrBadInput           = errors.New("Given parameter is not valid")
	ErrNoRecordFound      = errors.New("Record not Found")
	ErrMalformedParameter = errors.New("Malformed Parameter")
)
