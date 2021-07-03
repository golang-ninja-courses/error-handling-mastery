package pkga

type EOF struct{}

func (b EOF) Error() string {
	return "end of file"
}
