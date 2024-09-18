package derrors

type basicError struct {
	err  error
	code int
	kind ErrKind
}

func (e *basicError) Error() string {
	return e.err.Error()
}

func (e *basicError) Code() int {
	return e.code
}

func (e *basicError) Kind() ErrKind {
	return e.kind
}
