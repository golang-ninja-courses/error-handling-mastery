package db_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	db "github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/examples/05-errors-best-practices/api-borders"
)

func TestGetUserByIDOriginal(t *testing.T) {
	_, err := db.GetUserByIDOriginal(context.Background(), "uid")
	require.Error(t, err)
	assert.ErrorIs(t, err, sql.ErrNoRows) // Прыгаем через слои.
}

func TestGetUserByIDOwnError(t *testing.T) {
	_, err := db.GetUserByIDOwnError(context.Background(), "uid")
	require.Error(t, err)
	assert.ErrorIs(t, err, db.ErrNotFound)
	assert.False(t, errors.Is(err, sql.ErrNoRows))
}

func TestIsNotFoundError(t *testing.T) {
	_, err := db.GetUserByIDOriginal(context.Background(), "uid")
	require.Error(t, err)
	assert.True(t, db.IsNotFoundError(err))
}

func TestIsNotFoundErrorPrivate(t *testing.T) {
	_, err := db.GetUserByIDOwnPrivateError(context.Background(), "uid")
	require.Error(t, err)
	assert.True(t, db.IsNotFoundError2(err))
}
