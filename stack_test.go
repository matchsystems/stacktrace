package stacktrace_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/matchsystems/stacktrace"
)

func parentGetStacktrace() []stacktrace.Frame {
	return stacktrace.GetStacktrace(0, 10)
}

func nestedGetStacktrace() []stacktrace.Frame {
	return parentGetStacktrace()
}

type testStack struct {
	Function string
	Module   string
	File     string
	Line     int
}

func testFrames(t *testing.T, frames []stacktrace.Frame, testFrames []testStack) {
	t.Helper()

	for i, f := range frames {
		w := testFrames[i]
		require.Equal(t, w.Function, f.Function)
		require.Equal(t, w.Module, f.Module)
		require.Equal(t, w.Line, f.Line)
		require.Contains(t, f.File, w.File)
	}
}

func Test_GetStacktrace(t *testing.T) {
	t.Parallel()

	t.Run("single", func(t *testing.T) {
		t.Parallel()

		frames := stacktrace.GetStacktrace(0, 10)
		require.Len(t, frames, 2)

		testFrames(t, frames, []testStack{
			{
				Function: "GetStacktrace",
				Module:   "github.com/matchsystems/stacktrace",
				File:     "stack.go",
				Line:     10,
			},
			{
				Function: "Test_GetStacktrace.func1",
				Module:   "github.com/matchsystems/stacktrace_test",
				File:     "stack_test.go",
				Line:     44,
			},
		})
	})

	t.Run("parent", func(t *testing.T) {
		t.Parallel()

		frames := parentGetStacktrace()
		require.Len(t, frames, 3)

		testFrames(t, frames, []testStack{
			{
				Function: "GetStacktrace",
				Module:   "github.com/matchsystems/stacktrace",
				File:     "stack.go",
				Line:     10,
			},
			{
				Function: "parentGetStacktrace",
				Module:   "github.com/matchsystems/stacktrace_test",
				File:     "stack_test.go",
				Line:     12,
			},
			{
				Function: "Test_GetStacktrace.func2",
				Module:   "github.com/matchsystems/stacktrace_test",
				File:     "stack_test.go",
				Line:     66,
			},
		})
	})

	t.Run("nested", func(t *testing.T) {
		t.Parallel()

		frames := nestedGetStacktrace()
		require.Len(t, frames, 4)

		testFrames(t, frames, []testStack{
			{
				Function: "GetStacktrace",
				Module:   "github.com/matchsystems/stacktrace",
				File:     "stack.go",
				Line:     10,
			},
			{
				Function: "parentGetStacktrace",
				Module:   "github.com/matchsystems/stacktrace_test",
				File:     "stack_test.go",
				Line:     12,
			},
			{
				Function: "nestedGetStacktrace",
				Module:   "github.com/matchsystems/stacktrace_test",
				File:     "stack_test.go",
				Line:     16,
			},
			{
				Function: "Test_GetStacktrace.func3",
				Module:   "github.com/matchsystems/stacktrace_test",
				File:     "stack_test.go",
				Line:     94,
			},
		})
	})
}
