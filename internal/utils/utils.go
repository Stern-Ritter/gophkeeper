package utils

import (
	"fmt"
)

// Coalesce returns the first non-zero value from the two provided arguments.
// If the first value is non-zero, it will be returned.
// Otherwise, the second value is returned.
//
// This function supports the types string, int, and bool.
func Coalesce[T string | int | bool](firstValue, secondValue T) T {
	var zeroValue T
	if firstValue != zeroValue {
		return firstValue
	}
	return secondValue
}

// FormatBytes takes a number of bytes and returns a string representation
// of the size using the most appropriate unit of measurement (B, KB, MB, GB).
func FormatBytes(bytes int64) string {
	switch {
	case bytes < 1024:
		return fmt.Sprintf("%d B", bytes)
	case bytes < 1024*1024:
		return fmt.Sprintf("%.2f KB", float64(bytes)/1024)
	case bytes < 1024*1024*1024:
		return fmt.Sprintf("%.2f MB", float64(bytes)/(1024*1024))
	default:
		return fmt.Sprintf("%.2f GB", float64(bytes)/(1024*1024*1024))
	}
}
