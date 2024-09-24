package queue

import (
	"io"
	"time"
)

const defaultPostpone = time.Second

type AlreadyDoneError struct {}
func (e AlreadyDoneError) Error() string {
    return "job is already done"
}

type InconsistentDataError struct {}
func (e InconsistentDataError) Error() string {
	return "job payload is corrupted"
}

type InvalidIDError struct {}
func (e InvalidIDError) Error() string {
	return "invalid job id"
}

type NotFoundError struct {}
func (e NotFoundError) Error() string {
	return "job wasn't found"
}

type NotReadyError struct {}
func (e NotReadyError) Error() string {
	return "job is not ready to be performed"
}

var (
	ErrAlreadyDone      error = new(AlreadyDoneError)
	ErrInconsistentData error = new(InconsistentDataError)
	ErrInvalidID        error = new(InvalidIDError)
	ErrNotFound         error = new(NotFoundError)
	ErrNotReady         error = new(NotReadyError)
)

type Job struct {
	ID int
}

type Handler struct{}

func (h *Handler) Handle(job Job) (postpone time.Duration, err error) {
	err = h.process(job)
	if err != nil {
		if err == ErrAlreadyDone || err == ErrInconsistentData || err == ErrInvalidID || err == ErrNotFound {
			return 0, nil
		} else if err == ErrNotReady {
			return time.Second, nil
		} 
		return 0, err
	}
	return 0, nil
}

func (h *Handler) process(job Job) error {
	switch job.ID {
	case 1:
		return ErrInconsistentData
	case 2:
		return ErrNotReady
	case 3:
		return ErrNotFound
	case 4:
		return ErrAlreadyDone
	case 5:
		return ErrInvalidID
	case 6:
		return io.EOF
	}
	return nil
}
