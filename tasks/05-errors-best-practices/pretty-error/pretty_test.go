package errs

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPretty(t *testing.T) {
	err := fmt.Errorf("cannot get user schedule: %w",
		fmt.Errorf("cannot build data for event: %w",
			fmt.Errorf("cannot get image: %w",
				fmt.Errorf("cannot get image: %w",
					fmt.Errorf("cannot get image: %w",
						fmt.Errorf("cannot get image for event: %w", sql.ErrNoRows))))))
	t.Log(err)

	pErr := Pretty(err)
	require.Error(t, pErr)
	t.Log(pErr)

	assert.ErrorIs(t, pErr, sql.ErrNoRows)
	assert.Equal(t, `cannot get user schedule:
cannot build data for event:
cannot get image:
cannot get image:
cannot get image:
cannot get image for event:
`+sql.ErrNoRows.Error(), pErr.Error())
}

func TestPretty_NoError(t *testing.T) {
	err := Pretty(nil)
	require.NoError(t, err)
}
