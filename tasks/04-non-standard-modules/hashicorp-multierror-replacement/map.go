package multierror

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
)

// MapV2 работает как Map, только реализована на базе стандартной библиотеки.
func MapV2[T any](fn func(v T) error, input []T) error {
	// Реализуй меня.
	return nil
}

// Map применяет функцию fn к каждому элементу input.
// Если функция завершается ошибкой, то ошибка добавляется в результирующую ошибку, возвращаемую Map.
func Map[T any](fn func(v T) error, input []T) error {
	var result error
	for i, s := range input {
		if err := fn(s); err != nil {
			result = multierror.Append(result, fmt.Errorf("elem at %d: %w", i, err))
		}
	}
	return result
}
