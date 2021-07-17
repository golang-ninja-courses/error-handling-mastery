package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var errNotFound = errors.New("obj not found")

func GetUserByIDOwnPrivateError(ctx context.Context, uid UserID) (*User, error) {
	err := sql.ErrNoRows

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("exec query: %w: %v", errNotFound, err)
	}
	return &User{ID: "42"}, nil
}
