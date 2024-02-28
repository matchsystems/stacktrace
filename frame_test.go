package stacktrace_test

import (
	"testing"

	"github.com/matchsystems/stacktrace"
	"github.com/stretchr/testify/require"
)

func TestPrettyStack(t *testing.T) {
	t.Parallel()

	stack := []stacktrace.Frame{
		{"Func1", "module1", "file1.go", 10},
		{"Func2", "module2", "file2.go", 20},
	}

	expected := []string{"file1.go:10#Func1", "file2.go:20#Func2"}

	result := stacktrace.PrettyStack(stack)
	require.Equal(t, expected, result)
}
