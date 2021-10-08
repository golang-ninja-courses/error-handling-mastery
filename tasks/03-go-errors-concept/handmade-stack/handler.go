package handmadestack

import (
	"errors"
)

var (
	ErrExecSQL         = errors.New("exec sql error")
	ErrInitTransaction = errors.New("init transaction error")
)

type Entity struct {
	ID string
}

// Используются тестами.
var (
	getEntity        = func() (Entity, error) { return Entity{ID: "some-id"}, nil }
	updateEntity     = func(e Entity) error { return nil }
	runInTransaction = func(f func() error) error { return f() }
)

// Перепиши меня так, чтобы логика сохранилась,
// но путь до каждой ошибки был очевиден.
func handler() (Entity, error) {
	var e Entity

	if err := runInTransaction(func() (opErr error) {
		e, opErr = getEntity()
		if opErr != nil {
			return opErr
		}

		return updateEntity(e)
	}); err != nil {
		return Entity{}, err
	}

	if err := runInTransaction(func() error {
		return updateEntity(e)
	}); err != nil {
		return Entity{}, err
	}

	if err := runInTransaction(func() (opErr error) {
		return updateEntity(e)
	}); err != nil {
		return Entity{}, err
	}

	return e, nil
}
