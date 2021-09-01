package main

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type User struct {
	ID           string
	Email        string
	RegisteredAt time.Time
}

func TestAssertOnly_1(t *testing.T) {
	getUser := func() (*User, error) {
		return nil, errors.New("user not found")
	}

	u, err := getUser()
	assert.NoError(t, err)
	assert.Equal(t, "user-id", u.ID) // Получим панику "... nil pointer dereference".
	assert.Equal(t, "user-email", u.Email)
}

func TestAssertOnly_2(t *testing.T) {
	getUsers := func() ([]User, error) {
		return nil, nil
	}

	users, err := getUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 3)

	u := users[0] // Получим панику "... index out of range [0] ...".
	assert.Equal(t, "user-id", u.ID)
	assert.Equal(t, "user-email", u.Email)
}

func TestRequireAndAssert(t *testing.T) {
	getUser := func() (*User, error) {
		return new(User), nil
	}

	u, err := getUser()
	require.NoError(t, err)
	assert.Equal(t, "user-id", u.ID)
	assert.Equal(t, "user-email", u.Email)
}

func TestRequireOnly(t *testing.T) {
	getUser := func() (*User, error) {
		return new(User), nil
	}

	u, err := getUser()
	require.NoError(t, err)
	require.Equal(t, "user-id", u.ID) // Выйдем, например, уже здесь, хотя удобнее увидеть сразу все ошибки.
	require.Equal(t, "user-email", u.Email)
}
