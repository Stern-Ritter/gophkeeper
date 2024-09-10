package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileToFileMessage(t *testing.T) {
	file := File{
		ID:      "1",
		UserID:  "2",
		Name:    "file name",
		Size:    1024,
		Comment: "file comment",
	}

	message := FileToFileMessage(file)

	assert.Equal(t, "1", message.Id)
	assert.Equal(t, "2", message.UserId)
	assert.Equal(t, "file name", message.Name)
	assert.Equal(t, int64(1024), message.Size)
	assert.Equal(t, "file comment", message.Comment)
}

func TestFilesToRepeatedFileMessage(t *testing.T) {
	files := []File{
		{
			ID:      "1",
			UserID:  "1",
			Name:    "first file",
			Size:    2048,
			Comment: "first comment",
		},
		{
			ID:      "2",
			UserID:  "2",
			Name:    "second file",
			Size:    4096,
			Comment: "second comment",
		},
	}

	messages := FilesToRepeatedFileMessage(files)

	require.Len(t, messages, 2)

	first := messages[0]
	assert.Equal(t, "1", first.Id)
	assert.Equal(t, "1", first.UserId)
	assert.Equal(t, "first file", first.Name)
	assert.Equal(t, int64(2048), first.Size)
	assert.Equal(t, "first comment", first.Comment)

	second := messages[1]
	assert.Equal(t, "2", second.Id)
	assert.Equal(t, "2", second.UserId)
	assert.Equal(t, "second file", second.Name)
	assert.Equal(t, int64(4096), second.Size)
	assert.Equal(t, "second comment", second.Comment)
}
