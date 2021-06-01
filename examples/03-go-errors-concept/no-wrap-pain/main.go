package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrExecSQL         = errors.New("exec sql error")
	ErrInitTransaction = errors.New("init transaction error")
)

type Entity struct {
	ID string
}

func errOccurred() bool {
	return rand.Float64() < 0.5
}

func getEntity() (Entity, error) {
	if errOccurred() {
		return Entity{}, ErrExecSQL
	}
	return Entity{ID: "some-id"}, nil
}

func updateEntity(e Entity) error {
	if errOccurred() {
		return ErrExecSQL
	}
	return nil
}

func runInTransaction(f func() error) error {
	if errOccurred() {
		return ErrInitTransaction
	}
	return f()
}

func handler() error {
	var e Entity

	if err := runInTransaction(func() (opErr error) {
		e, opErr = getEntity()
		return opErr
	}); err != nil {
		return err
	}

	if err := runInTransaction(func() error {
		return updateEntity(e)
	}); err != nil {
		return err
	}

	if err := runInTransaction(func() (opErr error) {
		return updateEntity(e)
	}); err != nil {
		return err
	}

	return nil
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(handler())
	}
	/*
		От какой операции именно пришла ошибка?
			exec sql error
			init transaction error
			exec sql error
			exec sql error
			<nil>
	*/
}
