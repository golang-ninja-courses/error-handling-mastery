package allocator

import "fmt"

const (
	Admin          = 777
	MinMemoryBlock = 1024
)

type NotPermittedError struct{}

type ArgOutOfDomainError struct{}

func (npe *NotPermittedError) Error() string {
	return fmt.Sprintf("operation not permitted")
}

func (auod *ArgOutOfDomainError) Error() string {
	return fmt.Sprintf("numerical argument out of domain of func")
}

func Allocate(userID, size int) ([]byte, error) {
    if userID != Admin {
		return nil, new(NotPermittedError)
	}
	if size < MinMemoryBlock {
		return nil, new(ArgOutOfDomainError)
	}


	result := make([]byte, size, size)

	return result, nil
}
