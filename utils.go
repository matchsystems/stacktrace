package stacktrace

import (
	"runtime"
	"strings"
)

// packageName returns the package part of the symbol name, or the empty string
// if there is none.
// It replicates https://golang.org/pkg/debug/gosym/#Sym.PackageName, avoiding a
// dependency on debug/gosym.
func packageName(name string) string {
	// A prefix of "type." and "go." is a compiler-generated symbol that doesn't belong to any package.
	// See variable reservedimports in cmd/compile/internal/gc/subr.go
	if strings.HasPrefix(name, "go.") || strings.HasPrefix(name, "type.") {
		return ""
	}

	pathend := strings.LastIndex(name, "/")
	if pathend < 0 {
		pathend = 0
	}

	if i := strings.Index(name[pathend:], "."); i != -1 {
		return name[:pathend+i]
	}

	return ""
}

// splitQualifiedFunctionName splits a package path-qualified function name into
// package name and function name. Such qualified names are found in
// runtime.Frame.Function values.
func splitQualifiedFunctionName(name string) (string, string) {
	if name == "" {
		return "", ""
	}

	pkg := packageName(name)
	fun := strings.TrimPrefix(name, pkg+".")

	return pkg, fun
}

func ExtractFrames(pcs []uintptr) Frames {
	var frames Frames

	if len(pcs) == 0 {
		return frames
	}

	callersFrames := runtime.CallersFrames(pcs)

	for {
		callerFrame, more := callersFrames.Next()
		if len(callerFrame.File) > 0 {
			f := NewFrame(callerFrame)

			// Skip Go internal frames.
			if f.Module == "runtime" || f.Module == "testing" {
				if !more {
					break
				}

				continue
			}

			frames = append(frames, f)
		}

		if !more {
			break
		}
	}

	return frames
}
