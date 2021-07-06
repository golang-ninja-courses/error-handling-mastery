package docker

type Error struct {
}

func (e *Error) Error() string {
	panic("implement me")
}

func newDockerError(err error) *Error {
	return &Error{}
}
