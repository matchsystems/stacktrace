package stacktrace_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/matchsystems/stacktrace"
)

func c() []uintptr {
	pcs := make([]uintptr, 10)
	runtime.Callers(1, pcs)

	return pcs
}
func b() []uintptr { return c() }
func a() []uintptr { return b() }

func TestDefaultErrorStackMarshaler(t *testing.T) {
	t.Parallel()

	t.Run("single call", func(t *testing.T) {
		t.Parallel()

		callersFrames := stacktrace.ExtractFrames(c()).Pretty()

		require.Equal(t, []string{
			"stacktrace/utils_test.go:14#c",
			"stacktrace/utils_test.go:27#TestDefaultErrorStackMarshaler.func1",
		}, callersFrames)
	})

	t.Run("deep call", func(t *testing.T) {
		t.Parallel()

		callersFrames := stacktrace.ExtractFrames(a()).Pretty()

		require.Equal(t, []string{
			"stacktrace/utils_test.go:14#c",
			"stacktrace/utils_test.go:18#b",
			"stacktrace/utils_test.go:19#a",
			"stacktrace/utils_test.go:38#TestDefaultErrorStackMarshaler.func2",
		}, callersFrames)
	})

	t.Run("empty pcs", func(t *testing.T) {
		t.Parallel()

		callersFrames := stacktrace.ExtractFrames(nil).Pretty()

		require.Equal(t, []string{}, callersFrames)
	})

	t.Run("invalid pcs", func(t *testing.T) {
		t.Parallel()

		callersFrames := stacktrace.ExtractFrames([]uintptr{111}).Pretty()

		require.Equal(t, []string{}, callersFrames)
	})
}
