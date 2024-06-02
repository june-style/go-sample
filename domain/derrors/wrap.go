package derrors

import "fmt"

type wrapedError struct {
	err   error
	msg   string
	frame Frame
}

func (e *wrapedError) Unwrap() error {
	return e.err
}

func (e *wrapedError) Error() string {
	if e.msg == "" {
		return e.err.Error()
	}
	return fmt.Sprintf("%s: %s", e.err, e.msg)
}

func (e *wrapedError) Code() int {
	return Code(e.err)
}

func (e *wrapedError) Kind() ErrKind {
	return Kind(e.err)
}

func (e *wrapedError) Stacktrace() string {
	return e.frame.String()
}

type ErrParam struct {
	Key string
	Val any
}

func (e ErrParam) String() string {
	return fmt.Sprintf("%s -> %+v", e.Key, e.Val)
}
