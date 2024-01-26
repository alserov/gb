package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const testFilesFolder = "./test_files/"

func TestRun(t *testing.T) {
	tests := []struct {
		path   string
		exists bool
	}{
		{
			path:   "readme.md",
			exists: true,
		},
		{
			path:   "anything.json",
			exists: false,
		},
	}

	err := os.Mkdir(testFilesFolder, 0644)
	require.NoError(t, err)

	defer func() {
		err = os.RemoveAll(testFilesFolder)
		require.NoError(t, err)
	}()

	for _, tc := range tests {
		if tc.exists {
			f, err := os.OpenFile(testFilesFolder+tc.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			require.NoError(t, err)
			err = f.Close()
			require.NoError(t, err)
		}

		err = CheckFile(testFilesFolder + tc.path)
		require.Equal(t, tc.exists, err == nil)
	}
}
