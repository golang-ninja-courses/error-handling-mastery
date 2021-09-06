package example

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopyFileV3(t *testing.T) {
	dst := filepath.Join(os.TempDir(), "TestCopyFileV3")

	err := CopyFileV3("/etc/hosts", dst)
	require.NoError(t, err)

	defer func() {
		require.NoError(t, os.Remove(dst))
	}()

	{
		lhs, err := os.Stat("/etc/hosts")
		require.NoError(t, err)

		rhs, err := os.Stat(dst)
		require.NoError(t, err)

		require.Equal(t, lhs.Size(), rhs.Size())
	}
}
