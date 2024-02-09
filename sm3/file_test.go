package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCollection(t *testing.T) {
	c := NewCollection(FILE)

	t.Run("testing file collection implementation", func(t *testing.T) {
		err := c.Push("link.com", WithName("My link"), WithTags([]string{"tag1"}))
		require.NoError(t, err)

		err = c.Push("link.com", WithTags([]string{"tag2"}))
		require.NoError(t, err)

		count, err := c.Delete("", WithTag("tag1"))
		require.NoError(t, err)
		require.Equal(t, 1, count)

		links, err := c.Get("link.com")
		require.NoError(t, err)
		require.NotEmpty(t, links)

		links, err = c.All()
		require.NoError(t, err)
		require.Equal(t, c.Length(), len(links))

		err = c.Clear()
		require.NoError(t, err)

		links, err = c.All()
		require.NoError(t, err)
		require.Equal(t, c.Length(), len(links))
	})
}
