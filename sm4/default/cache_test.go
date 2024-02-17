package _default

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var c Cache

func init() {
	c = newCacheImpl(time.Second)
}

func TestCacheImpl(t *testing.T) {
	tests := []struct {
		name string

		opType string
		key    string
		val    []byte

		valExists bool
	}{
		{
			name:   "set",
			opType: "set",
			key:    "key",
			val:    []byte("val"),
		},
		{
			name:      "get existing",
			opType:    "get",
			key:       "key",
			val:       []byte("val"),
			valExists: true,
		},
		{
			name:   "get not existing",
			opType: "get",
			key:    "val1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.opType {
			case "set":
				c.Set(tc.key, tc.val)
			case "get":
				v, ok := c.Get(tc.key)
				switch tc.valExists {
				case true:
					require.Equal(t, true, ok)
					require.Equal(t, tc.val, v)
				default:
					require.Equal(t, false, ok)
					require.Empty(t, v)
				}
			}
		})
	}
}

func TestCacheImplWithClear(t *testing.T) {
	tests := []struct {
		name string

		opType string
		key    string
		val    []byte
	}{
		{
			name:   "set",
			opType: "set",
			key:    "key",
			val:    []byte("val"),
		},
		{
			name:   "set1",
			opType: "set",
			key:    "key1",
			val:    []byte("val"),
		},
		{
			name:   "get existing",
			opType: "get",
			key:    "key",
			val:    []byte("val"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.opType {
			case "set":
				c.Set(tc.key, tc.val)
				time.Sleep(time.Second)
			case "get":
				v, ok := c.Get(tc.key)
				require.Equal(t, false, ok)
				require.Empty(t, v)
			}
		})
	}
}

func TestCacheImpl_SetWriteCache(t *testing.T) {
	tests := []struct {
		name string

		opType string
		key    string
		val    []byte
	}{
		{
			name:   "set",
			opType: "set",
			key:    "key",
			val:    []byte("val"),
		},
		{
			name:   "get existing",
			opType: "get",
			key:    "key",
			val:    []byte("val"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.opType {
			case "set":
				c.SetWriteCache(tc.key, tc.val)
			case "get":
				vals := c.GetWriteCache(true)
				for k, v := range vals {
					require.Equal(t, tc.val, v)
					require.Equal(t, tc.key, k)
				}
			}
		})
	}
}

func TestCacheImpl_GetWriteCache(t *testing.T) {
	tests := []struct {
		name string

		opType string
		key    string
		val    []byte
	}{
		{
			name:   "set",
			opType: "set",
			key:    "key",
			val:    []byte("val"),
		},
		{
			name:   "get existing",
			opType: "get",
			key:    "key",
			val:    []byte("val"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.opType {
			case "set":
				c.SetWriteCache(tc.key, tc.val)
			case "get":
				vals := c.GetWriteCache(true)
				for k, v := range vals {
					require.Equal(t, tc.val, v)
					require.Equal(t, tc.key, k)
				}
				require.Empty(t, c.GetWriteCache(false))
			}
		})
	}
}
