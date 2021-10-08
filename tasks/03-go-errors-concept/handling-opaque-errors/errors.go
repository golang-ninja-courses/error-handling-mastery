package queue

type AlreadyDoneError struct{}

func (e *AlreadyDoneError) Error() string { return "job is already done" }
func (e *AlreadyDoneError) Skip() bool    { return true }

type InconsistentDataError struct{}

func (e *InconsistentDataError) Error() string { return "job payload is corrupted" }
func (e *InconsistentDataError) Skip() bool    { return true }

type InvalidIDError struct{}

func (e *InvalidIDError) Error() string { return "invalid job id" }
func (e *InvalidIDError) Skip() bool    { return true }

type NotFoundError struct{}

func (e *NotFoundError) Error() string { return "job wasn't found" }
func (e *NotFoundError) Skip() bool    { return true }

type NotReadyError struct{}

func (e *NotReadyError) Error() string   { return "job is not ready to be performed" }
func (e *NotReadyError) Temporary() bool { return true }
