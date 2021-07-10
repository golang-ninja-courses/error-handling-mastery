package wrapping_opaque_errors

type NetworkError struct{}

func (n *NetworkError) Error() string {
	return "network error"
}

func (n *NetworkError) IsTemporary() bool {
	return true
}

type WithMessageError struct {
	message string
	err     error
}

func (n *WithMessageError) Error() string {
	return "network error"
}

func (n *WithMessageError) Unwrap() error {
	return n.err
}

type isTemporary interface {
	IsTemporary() bool
}

func IsTemporary(err error) bool {
	// TODO реализуй меня
	return false
}
