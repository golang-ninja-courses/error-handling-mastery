package rest

func Handle() error {
	var err *HTTPError

	if err2 := usefulWork(); err2 != nil {
		err = ErrInternalServerError
	}
	return err
}

var usefulWork = func() error {
	return nil
}
