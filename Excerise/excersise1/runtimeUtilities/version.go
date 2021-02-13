package runtimeUtilities

import (
	"runtime"
)

// Goversion retruns the current running GO version
func Goversion() string {
	return runtime.Version()
}
