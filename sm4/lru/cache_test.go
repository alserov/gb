package lru

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const lruLimit = 3

func TestNewLRU(t *testing.T) {
	c := NewLRU(lruLimit)

	c.Set("key", []byte("v"))
	c.Set("key1", []byte("v1"))
	c.Set("key2", []byte("v2"))

	v, ok := c.Get("key")
	require.True(t, ok)
	require.Equal(t, []byte("v"), v)

	c.Set("key3", []byte("v3"))

	v, ok = c.Get("key")
	require.False(t, ok)
	require.Empty(t, v)
}
