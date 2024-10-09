//go:build cgo && !nolink_libwasmvm

package cosmwasm

import (
	"wasmvm.mleku.dev/internal/api"
)

func libwasmvmVersionImpl() (string, error) {
	return api.LibwasmvmVersion()
}
