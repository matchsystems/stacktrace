package stacktrace

import (
	"fmt"
)

var (
	LineFormatter = ShortLineFormatterFn //nolint:gochecknoglobals // for custom setting
)

func DefaultLineFormatterFn(file string, line int, funcName string) string {
	return fmt.Sprintf("%s:%d#%s", file, line, funcName)
}

func ShortCaller(file string) string {
	short := file
	z := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			z++
			if z > 1 {
				break
			}
		}
	}

	return short
}

func ShortLineFormatterFn(file string, line int, funcName string) string {
	return fmt.Sprintf("%s:%d#%s", ShortCaller(file), line, funcName)
}
