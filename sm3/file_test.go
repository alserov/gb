package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCollection(t *testing.T) {
	c := NewCollection(FILE)

	t.Run("testing file collection implementation", func(t *testing.T) {
		err := c.Push("link.com", WithName("My link"), WithTags([]string{"tag1", "tag2", "tag3"}))
		require.NoError(t, err)
		err = c.Push("link1.com", WithName("My link"), WithTags([]string{"tag1", "tag2"}))
		require.NoError(t, err)

		err = c.Push("link2.com", WithTags([]string{"tag2"}))
		require.NoError(t, err)

		count, err := c.Delete("", WithTag("tag2"), WithLimit(1), WithOffset(1))
		require.NoError(t, err)
		require.Equal(t, 1, count)

		err = c.Push("link.com", WithName("My link"), WithTags([]string{"tag1", "tag2", "tag3"}))
		require.NoError(t, err)

		links, err := c.Get("link.com", WithOffset(1))
		require.NoError(t, err)
		require.Equal(t, 1, len(links))

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
