package stacktrace_test

import (
	"fmt"
	"testing"

	"github.com/matchsystems/stacktrace"
	"github.com/stretchr/testify/require"
)

func TestShortCaller(t *testing.T) {
	t.Parallel()

	require.Equal(t, "file.go", stacktrace.ShortCaller("file.go"))
	require.Equal(t, "file.go", stacktrace.ShortCaller("to/file.go"))
	require.Equal(t, "file.go", stacktrace.ShortCaller("///path//to///file.go"))
	require.Equal(t, "", stacktrace.ShortCaller(""))
	require.Equal(t, "file", stacktrace.ShortCaller("/path/to/file"))
}

func TestShortLineFormatterFn(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		file     string
		line     int
		funcName string
		expected string
	}{
		{"file.go", 10, "TestFunction", "file.go:10#TestFunction"},
		{"/path/to/file.go", 20, "AnotherFunction", "to/file.go:20#AnotherFunction"},
		{"", 0, "", ":0#"},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(fmt.Sprintf("%s:%d#%s", tc.file, tc.line, tc.funcName), func(t *testing.T) {
			t.Parallel()

			result := stacktrace.ShortLineFormatterFn(tc.file, tc.line, tc.funcName)
			require.Equal(t, tc.expected, result)
		})
	}
}
