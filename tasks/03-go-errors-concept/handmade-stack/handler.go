package handmadestack

import (
	"errors"
	"fmt"
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
			return fmt.Errorf("cannot get entity %v: %v", e, opErr)
		}

		upE1 := updateEntity(e)
		if upE1 != nil {
			return fmt.Errorf("can not update entity %v: %v", e, upE1)
		}
		return
	}); err != nil {
		return Entity{}, fmt.Errorf("transaction 1 error %v", err)
	}

	if err := runInTransaction(func() error {
        upE2 := updateEntity(e)
		if upE2 != nil {
			return fmt.Errorf("can not update entity %v: %v", e, upE2)
		}
		return nil
	}); err != nil {
		return Entity{}, fmt.Errorf("transaction 2 erorr %v", err)
	}

	if err := runInTransaction(func() (opErr error) {
        upE3 := updateEntity(e)
		if upE3 != nil {
			return fmt.Errorf("can not update entity %v: %v", e, upE3)
		}
		return nil
	}); err != nil {
		return Entity{}, fmt.Errorf("transaction 3 erorr %v", err)
	}

	return e, nil
}
