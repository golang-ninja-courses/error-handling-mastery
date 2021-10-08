package db

import (
	"database/sql"
	"errors"
)

func IsNotFoundError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func IsNotFoundError2(err error) bool {
	return errors.Is(err, errNotFound)
}
