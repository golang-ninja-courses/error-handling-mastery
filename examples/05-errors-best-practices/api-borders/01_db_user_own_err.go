package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("obj not found")

func GetUserByIDOwnError(ctx context.Context, uid UserID) (*User, error) {
	if err := selectUser(); err == sql.ErrNoRows {
		return nil, fmt.Errorf("exec query: %w: %v", ErrNotFound, err)
	}
	return &User{ID: "42"}, nil
}
