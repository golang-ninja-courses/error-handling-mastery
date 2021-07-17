package db

import (
	"context"
	"database/sql"
	"fmt"
)

type UserID string

type User struct {
	ID UserID
}

func GetUserByIDOriginal(ctx context.Context, uid UserID) (*User, error) {
	err := sql.ErrNoRows

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("exec query: %w", err)
	}
	return &User{ID: "42"}, nil
}
