package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	require.NoError(t, validate("abc.DEF-ghi_123"))
	require.Error(t, validate(""))
	require.Error(t, validate("."))
	require.Error(t, validate("\u65E5"))
	require.Error(t, validate(".a"))
	require.Error(t, validate("a\u65E5"))
}
