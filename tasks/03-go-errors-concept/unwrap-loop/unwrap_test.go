package errs

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type FileLoadError struct {
	URL string
	Err error
}

func (p *FileLoadError) Error() string {
	return fmt.Sprintf("%q: %v", p.URL, p.Err)
}

func (p *FileLoadError) Unwrap() error {
	return p.Err
}

func TestUnwrap(t *testing.T) {
	cases := []struct {
		name     string
		inputErr error
		expected error
	}{
		{
			name: "FileLoadError:FileLoadError:FileLoadError:FileLoadError:os.PathError:os.Permission",
			inputErr: &FileLoadError{
				URL: "url1",
				Err: &FileLoadError{
					URL: "url2",
					Err: &FileLoadError{
						URL: "url3",
						Err: &FileLoadError{
							URL: "url4",
							Err: &os.PathError{
								Path: "path",
								Err:  os.ErrPermission,
							},
						},
					},
				},
			},
			expected: os.ErrPermission,
		},
		{
			name: "FileLoadError:FileLoadError:FileLoadError:nil",
			inputErr: &FileLoadError{
				URL: "url1",
				Err: &FileLoadError{
					URL: "url2",
					Err: &FileLoadError{
						URL: "url3",
						Err: nil,
					},
				},
			},
			expected: nil,
		},
		{
			name:     "nil",
			inputErr: nil,
			expected: nil,
		},
		{
			name:     "io.EOF",
			inputErr: io.EOF,
			expected: io.EOF,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := Unwrap(tt.inputErr)
			// NOTE: В данном случае использование require.ErrorIs некорректно.
			require.Equal(t, tt.expected, err)
		})
	}
}
