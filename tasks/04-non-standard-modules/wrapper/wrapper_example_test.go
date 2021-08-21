package errors

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

func ExampleNewUserError() {
	err := errors.Wrap(NewUserError(io.EOF, "Bob"), "message")

	type withUserID interface {
		UserID() string
	}

	var i withUserID
	if errors.As(err, &i) {
		fmt.Println(i.UserID())
	}

	if i, ok := errors.Cause(err).(withUserID); ok { // Это не работает!
		fmt.Println(i.UserID())
	}

	// Output:
	// Bob
}

func NewUserError(err error, userID string) error {
	return &userError{
		err:    err,
		userID: userID,
	}
}

var _ Wrapper = (*userError)(nil)

type userError struct {
	err    error
	userID string
}

func (ie *userError) Error() string {
	return fmt.Sprintf("user %s: %v", ie.userID, ie.err)
}

func (ie *userError) Cause() error {
	return ie.err
}

func (ie *userError) Unwrap() error {
	return ie.err
}

func (ie *userError) UserID() string {
	return ie.userID
}
