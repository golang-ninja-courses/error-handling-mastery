package allocator

const (
	Admin          = 777
	MinMemoryBlock = 1024
)

type NotPermittedError struct{}

type ArgOutOfDomainError struct{}

func Allocate(userID, size int) ([]byte, error) {
	// Реализуй меня.
	return nil, nil
}
