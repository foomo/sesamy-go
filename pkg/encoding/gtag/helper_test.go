package gtag_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

// readBody returns the encoded body, requiring it to be non-nil.
func readBody(t *testing.T, body io.Reader) string {
	t.Helper()
	require.NotNil(t, body)
	out, err := io.ReadAll(body)
	require.NoError(t, err)
	return string(out)
}
