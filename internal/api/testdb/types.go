package testdb

import (
	"errors"

	"wasmvm.mleku.dev/types"
)

var (

	// errKeyEmpty is returned when attempting to use an empty or nil key.
	errKeyEmpty = errors.New("key cannot be empty")

	// errValueNil is returned when attempting to set a nil value.
	errValueNil = errors.New("value cannot be nil")
)

type Iterator = types.Iterator
