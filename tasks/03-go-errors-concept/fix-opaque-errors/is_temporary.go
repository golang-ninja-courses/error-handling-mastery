package errors

func IsTemporary(err error) bool {
	type t interface {
		IsTemporary() bool
	}
	// Реализуй меня.
	return false
}
