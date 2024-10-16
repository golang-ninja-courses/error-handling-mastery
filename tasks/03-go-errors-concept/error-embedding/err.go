package errors


type Err struct {
	Msg string
}

type AlreadyDoneError struct {
	coreErr Err
}
func (e *AlreadyDoneError) Error() string { return "job is already done" }

type InconsistentDataError struct {
	coreErr Err
}
func (e *InconsistentDataError) Error() string { return "job payload is corrupted" }

type InvalidIDError struct {
	coreErr Err
}
func (e *InvalidIDError) Error() string { return "invalid job id" }

type NotReadyError struct {
	coreErr Err
}
func (e *NotFoundError) Error() string { return "job wasn't found"}

type NotFoundError struct {
	coreErr Err
}
func (e *NotReadyError) Error() string { return "job is not ready to be performed"}

var (
	ErrAlreadyDone      error = &AlreadyDoneError{Err{"job is already done"}}
	ErrInconsistentData error = &InconsistentDataError{Err{"job payload is corrupted"}}
	ErrInvalidID        error = &InvalidIDError{Err{"invalid job id"}}
	ErrNotReady         error = &NotReadyError{Err{"job is not ready to be performed"}}
	ErrNotFound         error = &NotFoundError{Err{"job wasn't found"}}
)


// Реализуй тип Err и типы для ошибок выше, используя его.
