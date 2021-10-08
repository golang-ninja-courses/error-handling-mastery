package queue

import (
	"io"
	"time"
)

const defaultPostpone = time.Second

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
		// Обработайте ошибку.
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
