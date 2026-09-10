package mpv2_test

import (
	"encoding/json"
	"testing"

	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserData_OmitsEmptyAddress(t *testing.T) {
	t.Parallel()

	v := mpv2.UserData{
		SHA256EmailAddress: mpv2.NewSHA256Hash("foo@example.com"),
	}

	out, err := json.Marshal(v)
	require.NoError(t, err)

	// `omitempty` never omits a struct, so this used to emit `"address":{}`.
	// Match on the key itself; "sha256_email_address" also contains "address".
	assert.NotContains(t, string(out), `"address"`)

	var got map[string]any
	require.NoError(t, json.Unmarshal(out, &got))
	assert.NotContains(t, got, "address")
}

func TestUserData_KeepsPopulatedAddress(t *testing.T) {
	t.Parallel()

	v := mpv2.UserData{
		Address: mpv2.UserDataAddress{City: "Munich"},
	}

	out, err := json.Marshal(v)
	require.NoError(t, err)

	assert.JSONEq(t, `{"address":{"city":"Munich"}}`, string(out))
}
