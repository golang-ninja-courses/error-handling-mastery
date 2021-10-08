package errors

var (
	ErrAlreadyDone      error = &AlreadyDoneError{Err{"job is already done"}}
	ErrInconsistentData error = &InconsistentDataError{Err{"job payload is corrupted"}}
	ErrInvalidID        error = &InvalidIDError{Err{"invalid job id"}}
	ErrNotReady         error = &NotReadyError{Err{"job is not ready to be performed"}}
	ErrNotFound         error = &NotFoundError{Err{"job wasn't found"}}
)

// Реализуй тип Err и типы для ошибок выше, используя его.
