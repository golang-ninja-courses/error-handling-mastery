package main

import "fmt"

var (
	ErrInconsistentData error = &InconsistentDataError{Err{"job payload is corrupted"}}
	ErrNotReady         error = &NotReadyError{Err{"job is not ready to be performed"}}
	ErrNotFound         error = &NotFoundError{Err{"job wasn't found"}}
	ErrAlreadyDone      error = &AlreadyDoneError{Err{"job is already done"}}
	ErrInvalidID        error = &InvalidIDError{Err{"invalid job id"}}
)

type Err struct {
	msg string
}

func (e Err) Error() string {
	return e.msg
}

type InconsistentDataError struct {
	Err
}

type NotReadyError struct {
	Err
}

type NotFoundError struct {
	Err
}

type AlreadyDoneError struct {
	Err
}

type InvalidIDError struct {
	Err
}

func main() {
	for _, e := range []error{
		ErrInconsistentData,
		ErrNotReady,
		ErrNotFound,
		ErrAlreadyDone,
		ErrInvalidID,
	} {
		fmt.Println(e)
	}
}
