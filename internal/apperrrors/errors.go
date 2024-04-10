package apperrrors

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrScanFailed       = errors.New("database read failed")
	ErrParsingPathParam = errors.New("error parsing path param")
	ErrInvalidPathParam = errors.New("invalid path param")
)
