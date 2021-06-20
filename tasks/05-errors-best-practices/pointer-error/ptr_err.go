package ptrerror

type PointerError struct {
	msg string
}

func (e *PointerError) Error() string {
	return e.msg
}

func NewPointerError(m string) PointerError {
	return PointerError{msg: m}
}
