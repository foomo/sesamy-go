package session_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/foomo/sesamy-go/pkg/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGARegex(t *testing.T) {
	t.Parallel()

	ga1 := "GA1.1.584335997.1746564151"

	{
		require.True(t, session.GA1Regex.MatchString(ga1))
		matches := session.GA1Regex.FindStringSubmatch(ga1)
		require.Len(t, matches, 4)
		assert.Equal(t, "1", matches[1])
		assert.Equal(t, "584335997", matches[2])
		assert.Equal(t, "1746564151", matches[3])
	}
}

func TestGSRegex(t *testing.T) {
	t.Parallel()

	gs1 := "GS1.1.1746524658.33.0.1746524658.0.0.0"
	gs2 := "GS2.1.s1746524658$o1$g0$t1746562809$j0$l0$h596732783"

	{
		require.True(t, session.GS1Regex.MatchString(gs1))
		matches := session.GS1Regex.FindStringSubmatch(gs1)
		require.Len(t, matches, 9)
		assert.Equal(t, "1", matches[1])
		assert.Equal(t, "1746524658", matches[2])
		assert.Equal(t, "33", matches[3])
		assert.Equal(t, "0", matches[4])
		assert.Equal(t, "1746524658", matches[5])
		assert.Equal(t, "0", matches[6])
		assert.Equal(t, "0", matches[7])
		assert.Equal(t, "0", matches[8])

		assert.False(t, session.GS1Regex.MatchString(gs2))
		assert.Empty(t, session.GS1Regex.FindStringSubmatch(gs2))
	}

	{
		require.True(t, session.GS2Regex.MatchString(gs2))
		matches := session.GS2Regex.FindStringSubmatch(gs2)
		require.Len(t, matches, 9)
		assert.Equal(t, "1", matches[1])
		assert.Equal(t, "1746524658", matches[2])
		assert.Equal(t, "1", matches[3])
		assert.Equal(t, "0", matches[4])
		assert.Equal(t, "1746562809", matches[5])
		assert.Equal(t, "0", matches[6])
		assert.Equal(t, "0", matches[7])
		assert.Equal(t, "596732783", matches[8])

		assert.False(t, session.GS2Regex.MatchString(gs1))
		assert.Empty(t, session.GS2Regex.FindStringSubmatch(gs1))
	}
}

func TestMeasurementIDSuffix(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "F9XM71K45T", session.MeasurementIDSuffix("G-F9XM71K45T"))
	assert.Equal(t, "F9XM71K45T", session.MeasurementIDSuffix("F9XM71K45T"))
}

func TestParseGA_NoCookie(t *testing.T) {
	t.Parallel()

	r := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)

	_, err := session.ParseGAClientID(r)
	require.ErrorIs(t, err, http.ErrNoCookie)

	_, err = session.ParseGASessionID(r, "F9XM71K45T")
	require.ErrorIs(t, err, http.ErrNoCookie)
}
