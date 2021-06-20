package queue

import (
	"io"
	"time"
)

const (
	defaultPostpone = time.Second
)

var (
	ErrInconsistentData error
	ErrNotReady         error
	ErrNotFound         error
	ErrAlreadyDone      error
	ErrInvalidID        error
)

type Job struct {
	ID int `json:"id"`
}

type Handler struct{}

func (j *Handler) Handle(job Job) (postpone int64, err error) {
	err = j.process(job)
	if err != nil {
		// Обработайте ошибку.
	}

	return 0, nil
}

func (j *Handler) process(job Job) error {
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
