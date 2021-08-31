package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	listener, err := createEventListener()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := listener(); err != nil {
		fmt.Println(err)
	}
}

func createEventListener() (func() error, error) {
	obj, err := getObject()
	if err != nil {
		return nil, errors.Wrap(err, "get object")
	}

	if obj.ID == 42 {
		return nil, errors.Wrap(err, "events is not supported for this obj")
	}

	return func() error {
		fmt.Println("ok")
		return nil
	}, nil
}

type Object struct {
	ID int
}

func getObject() (*Object, error) {
	return &Object{ID: 42}, nil
}
