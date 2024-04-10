package apperrrors

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrScanFailed       = errors.New("registe failed")
	ErrParsingPathParam = errors.New("error parsing path param")
	ErrInvalidPathParam = errors.New("invalid path param")
)
