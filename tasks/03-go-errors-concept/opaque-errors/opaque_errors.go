package main

import (
	"io"
)

type shouldBeSkipped interface {
	Skip() bool
}

type temporary interface {
	Temporary() bool
}

type InconsistentDataError struct{}

func (e *InconsistentDataError) Error() string {
	return "job payload is corrupted"
}

type NotReadyError struct{}

func (e *NotReadyError) Error() string {
	return "job is not ready yet"
}

type Job struct {
	ID int `json:"id"`
}

type Handler struct{}

func (j *Handler) process(job Job) error {
	switch job.ID {
	case 1:
		return &NotReadyError{}
	case 2:
		return &InconsistentDataError{}
	case 3:
		return nil
	case 4:
		return io.EOF
	}
	return nil
}

func (j *Handler) Handle(job Job) (postpone int64, err error) {
	err = j.process(job)
	if err != nil {
		// Обработайте ошибку.
	}

	return 0, err
}
