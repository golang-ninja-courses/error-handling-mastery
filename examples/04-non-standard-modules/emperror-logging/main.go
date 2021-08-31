package main

import (
	"errors"

	emperrors "emperror.dev/errors"
	logrushandler "emperror.dev/handler/logrus"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	handler := logrushandler.New(logger)

	err := errors.New("an error")
	handler.Handle(err)

	err2 := emperrors.WithDetails(err, "userID", 3587, "requestID", "4cfdc2e157eefe6facb983b1d557b3a1")
	handler.Handle(err2)
}
