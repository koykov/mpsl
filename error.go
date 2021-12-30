package mpsl

import "errors"

var (
	ErrBadDB    = errors.New("cache uninitialized, use New()")
	ErrNoHasher = errors.New("no hasher provided")
)
