package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatBytes(t *testing.T) {
	testCases := []struct {
		input int64
		want  string
	}{
		{512, "512 B"},
		{1023, "1023 B"},
		{1024, "1.00 KB"},
		{1136, "1.11 KB"},
		{1048576, "1.00 MB"},
		{1279262, "1.22 MB"},
		{1073741824, "1.00 GB"},
		{1428076625, "1.33 GB"},
	}

	for _, tt := range testCases {
		got := FormatBytes(tt.input)
		assert.Equal(t, tt.want, got,
			"Expected format for %d bytes to be '%s', but got '%s' instead", tt.input, tt.want, got)
	}
}

func TestCoalesce(t *testing.T) {
	t.Run("int values", func(t *testing.T) {
		result := Coalesce(5, 10)
		assert.Equal(t, 5, result,
			"Expected Coalesce(5, 10) to return 5, but got %d", result)

		result = Coalesce(0, 10)
		assert.Equal(t, 10, result,
			"Expected Coalesce(0, 10) to return 10, but got %d", result)
	})

	t.Run("string values", func(t *testing.T) {
		result := Coalesce("hello", "world")
		assert.Equal(t, "hello", result,
			"Expected Coalesce('hello', 'world') to return 'hello', but got '%s'", result)

		result = Coalesce("", "world")
		assert.Equal(t, "world", result,
			"Expected Coalesce('', 'world') to return 'world', but got '%s'", result)
	})

	t.Run("bool values", func(t *testing.T) {
		result := Coalesce(true, false)
		assert.Equal(t, true, result,
			"Expected Coalesce(true, false) to return true, but got %v", result)

		result = Coalesce(false, true)
		assert.Equal(t, true, result,
			"Expected Coalesce(false, true) to return true, but got %v", result)
	})

	t.Run("zero values", func(t *testing.T) {
		intResult := Coalesce(0, 0)
		assert.Equal(t, 0, intResult,
			"Expected Coalesce(0, 0) to return 0, but got %d", intResult)

		stringResult := Coalesce("", "")
		assert.Equal(t, "", stringResult,
			"Expected Coalesce('', '') to return an empty string, but got '%s'", stringResult)

		boolResult := Coalesce(false, false)
		assert.Equal(t, false, boolResult,
			"Expected Coalesce(false, false) to return false, but got %v", boolResult)
	})
}
