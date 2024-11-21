package sesamy_test

import (
	"testing"

	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeParams(t *testing.T) {
	type params struct {
		Title string `json:"title"`
	}

	event := sesamy.Event[any]{
		Name: "test",
		Params: map[string]any{
			"title":       "foo",
			"description": "foo",
		},
	}

	var p params
	require.NoError(t, event.DecodeParams(&p))
	assert.Equal(t, "foo", p.Title)

	p.Title = "bar"

	require.NoError(t, event.EncodeParams(p))
	assert.Equal(t, map[string]any{
		"title":       "bar",
		"description": "foo",
	}, event.Params)
}
