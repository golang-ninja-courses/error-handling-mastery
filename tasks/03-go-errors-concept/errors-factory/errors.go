package errors

// NewError возвращает новое значение-ошибку, текст которой является msg.
// Две ошибки с одинаковым текстом, созданные через NewError, не равны между собой:
//
//	NewError("end of file") != NewError("end of file")

type Err struct {
	Message string
}

func (e *Err) Error() string {
	return e.Message
} 

func NewError(msg string) error {
    e := &Err{
		Message: msg,
	}	
	return e
}
