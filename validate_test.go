package klevs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	require.NoError(t, ValidateName("abc.DEF-ghi_123"))
	require.NoError(t, ValidateName("."))
	require.NoError(t, ValidateName(".a"))
	require.NoError(t, ValidateName("__"))
	require.Error(t, ValidateName(""))
	require.Error(t, ValidateName("&"))
	require.Error(t, ValidateName("\u65E5"))
	require.Error(t, ValidateName("a\u65E5"))
}
