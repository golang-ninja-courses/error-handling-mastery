package queue

type AlreadyDoneError struct{}

func (e *AlreadyDoneError) Error() string { return "job is already done" }

type InconsistentDataError struct{}

func (e *InconsistentDataError) Error() string { return "job payload is corrupted" }

type InvalidIDError struct{}

func (e *InvalidIDError) Error() string { return "invalid job id" }

type NotFoundError struct{}

func (e *NotFoundError) Error() string { return "job wasn't found" }

type NotReadyError struct{}

func (e *NotReadyError) Error() string { return "job is not ready to be performed" }
