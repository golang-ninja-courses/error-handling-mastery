package main

import (
	"io"
)

type InconsistentDataError struct{}

func (e *InconsistentDataError) Error() string {
	return "job payload is corrupted"
}

type NotReadyError struct{}

func (e *NotReadyError) Error() string {
	return "job is not ready yet"
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "job wasn't found"
}

type AlreadyDoneError struct{}

func (e *AlreadyDoneError) Error() string {
	return "job is already done"
}

type InvalidIdError struct{}

func (e *InvalidIdError) Error() string {
	return "invalid job id"
}

type Job struct {
	ID int `json:"id"`
}

type Handler struct{}

func (j *Handler) process(job Job) error {
	switch job.ID {
	case 1:
		return &InconsistentDataError{}
	case 2:
		return &NotReadyError{}
	case 3:
		return &NotFoundError{}
	case 4:
		return &AlreadyDoneError{}
	case 5:
		return &InvalidIdError{}
	case 6:
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
