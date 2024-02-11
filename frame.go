package stacktrace

import (
	"runtime"
)

// Frame represents a function call and it's metadata.
type Frame struct {
	Function string `json:"func,omitempty"`
	Module   string `json:"module,omitempty"`
	File     string `json:"source,omitempty"`
	Line     int    `json:"line,string,omitempty"`
}

// NewFrame assembles a stacktrace frame out of runtime.Frame.
func NewFrame(f runtime.Frame) Frame {
	pkg, function := splitQualifiedFunctionName(f.Function)

	return Frame{
		File:     f.File,
		Line:     f.Line,
		Module:   pkg,
		Function: function,
	}
}

type Frames []Frame

func (f Frames) Pretty() []string {
	return PrettyStack(f)
}

func PrettyStack(stack []Frame) []string {
	out := make([]string, len(stack))

	for i := range stack {
		out[i] = LineFormatter(stack[i].File, stack[i].Line, stack[i].Function)
	}

	return out
}
