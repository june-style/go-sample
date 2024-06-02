package derrors

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Frame struct {
	file           string
	line           int
	programCounter uintptr
}

func NewFrame(file string, line int, programCounter uintptr) Frame {
	return Frame{
		file:           file,
		line:           line,
		programCounter: programCounter,
	}
}

func (f Frame) String() string {
	return fmt.Sprintf("%s:%d", f.file, f.line)
}

func (f Frame) Dir() string {
	return filepath.Dir(f.file)
}

func (f Frame) Func() string {
	if f.programCounter == 0 {
		return ""
	}
	fn := runtime.FuncForPC(f.programCounter)
	return fmt.Sprintf("%s()", fn.Name())
}
