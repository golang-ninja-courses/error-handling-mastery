package queue

import (
	"io"
	"time"
)

const defaultPostpone = time.Second

type Job struct {
	ID int
}

type Handler struct{}

func (h *Handler) Handle(job Job) (postpone time.Duration, err error) {
	err = h.process(job)
	if err != nil {
		if isTemporary(err) {
			return defaultPostpone, nil
		} else if shouldBeSkipped(err) {
			return 0, nil
		}
		return 0, err
	}

	return 0, nil
}

func isTemporary(err error) bool {
    e, ok := err.(interface{ Temporary() bool })
    return ok && e.Temporary()
}

func shouldBeSkipped(err error) bool {
	e, ok := err.(interface{ Skip() bool })
    return ok && e.Skip()
}

func (h *Handler) process(job Job) error {
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
		return &InvalidIDError{}
	case 6:
		return io.EOF
	}
	return nil
}
