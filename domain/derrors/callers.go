package derrors

import "runtime"

func Caller(skip int) Frame {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return Frame{}
	}
	return NewFrame(file, line, pc)
}

func Callers() (frames []Frame) {
	const depth = 32
	const skip = 4
	pcs := make([]uintptr, depth)
	cnt := runtime.Callers(skip, pcs)
	cfs := runtime.CallersFrames(pcs[:cnt])
	for {
		frame, hasNext := cfs.Next()
		frames = append(frames, NewFrame(
			frame.File, frame.Line, frame.PC,
		))
		if !hasNext {
			break
		}
	}
	return
}
