package derrors

import "slices"

func Unwrap(err error) error {
	type iface interface {
		Unwrap() error
	}

	if eErr, ok := err.(iface); ok {
		return eErr.Unwrap()
	}
	return nil
}

func Code(err error) int {
	type iface interface {
		Code() int
	}

	var code int
	if eErr, ok := err.(iface); ok {
		code = eErr.Code()
	}
	return code
}

func Kind(err error) ErrKind {
	type iface interface {
		Kind() ErrKind
	}

	kind := ErrKindUnknown
	if eErr, ok := err.(iface); ok {
		kind = eErr.Kind()
	}
	return kind
}

func Stacktrace(err error) []string {
	type iface interface {
		Stacktrace() string
	}

	var stacktrace []string
	for {
		if eErr, ok := err.(iface); ok {
			stacktrace = append(stacktrace, eErr.Stacktrace())
		}
		if err = Unwrap(err); err == nil {
			slices.Reverse(stacktrace)
			return stacktrace
		}
	}
}

func StacktraceAll() []string {
	var stacktrace []string
	for _, frame := range Callers() {
		stacktrace = append(stacktrace, frame.String())
	}
	return stacktrace
}
