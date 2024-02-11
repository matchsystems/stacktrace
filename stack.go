package stacktrace

import (
	"runtime"
)

// GetStacktrace creates a stacktrace using runtime.Callers.
func GetStacktrace(skip, limit int) Frames {
	pcs := make([]uintptr, limit)
	n := runtime.Callers(skip, pcs)

	if n == 0 {
		return nil
	}

	return ExtractFrames(pcs[:n])
}
