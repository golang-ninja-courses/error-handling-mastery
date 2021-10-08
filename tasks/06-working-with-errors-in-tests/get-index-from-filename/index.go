package index

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const prefix = "parsed_page_" // parsed_page_100

var (
	ErrInvalidFilename     = errors.New("invalid filename")
	ErrIndexMustBePositive = errors.New("index must be > 0")
)

func GetIndexFromFileName(fileName string) (int, error) {
	parts := strings.Split(fileName, prefix)
	if len(parts) != 2 || parts[1] == "" {
		return 0, fmt.Errorf("%w: no index in filename %q", ErrInvalidFilename, fileName)
	}

	num := parts[1]

	index, err := strconv.ParseInt(num, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("cannot parse index as int: %w", err)
	}

	if index <= 0 {
		return 0, fmt.Errorf("%w: got %d", ErrIndexMustBePositive, index)
	}

	return int(index), nil
}
