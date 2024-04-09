package apperrrors

import "errors"

var (
	ErrNotFound   = errors.New("not found")
	ErrScanFailed = errors.New("registe failed")
)
