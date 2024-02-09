package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewCollection(LINKED_LIST)

	t.Run("testing linked list collection implementation", func(t *testing.T) {
		l.Push("link1", WithName("My First Link"), WithTags([]string{"tag"}))

		count, err := l.Delete("", WithTag("tag"))
		require.NoError(t, err)
		require.Equal(t, 1, count)

		links, err := l.Get("link1")
		require.Empty(t, links)

		l.Push("link2", WithName("My Second Link"))
		l.Push("link2", WithName("My Third Link"))

		count, err = l.Delete("link2", WithLimit(1))
		require.Equal(t, 1, count)

		links, err = l.Get("link2")
		require.NoError(t, err)
		require.Equal(t, 1, len(links))

		l.Push("link4", WithName("My Fourth Link"), WithTags([]string{"tag"}))
		l.Push("link4", WithName("My Fifth Link"))
		l.Push("link4")

		count, err = l.Delete("link4", WithOffset(1), WithLimit(1))
		require.NoError(t, err)
		require.Equal(t, 1, count)

		links, err = l.All()
		require.NoError(t, err)

		require.Equal(t, l.Length(), len(links))

		l.Clear()
		links, err = l.All()
		require.NoError(t, err)
		require.Equal(t, l.Length(), len(links))
	})
}
