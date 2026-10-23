package resolve

import "errors"

var (
	// ErrNotFound means a source does not have the document or version (404, 410,
	// a missing file). Resolution moves on to the next source.
	ErrNotFound = errors.New("geul/resolve: not found")
	// ErrVerify means fetched bytes do not match the version list or the Key.
	// Resolution never falls through to another source on it.
	ErrVerify = errors.New("geul/resolve: verification failed")
)
