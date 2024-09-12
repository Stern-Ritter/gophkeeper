package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestInitialize_Success(t *testing.T) {
	for _, level := range []string{"debug", "info", "warn", "error", "dpanic", "panic", "fatal"} {
		t.Run(level, func(t *testing.T) {
			logger, err := Initialize(level)
			require.NoError(t, err, "Expected Initialize to succeed with valid logging level")
			require.NotNil(t, logger, "Expected Initialize to return a non-nil logger")

			expectedLevel, err := zap.ParseAtomicLevel(level)
			require.NoError(t, err, "Expected ParseAtomicLevel to succeed")
			assert.Equal(t, expectedLevel.Level(), logger.Level(), "Expected logger level to match the specified level")
		})
	}
}

func TestInitialize_InvalidLevel(t *testing.T) {
	logger, err := Initialize("invalid")
	require.Error(t, err, "Expected Initialize to return an error for an invalid logging level")
	require.Nil(t, logger, "Expected Initialize to return nil logger for an invalid logging level")
}

func TestInitialize_EmptyLevel(t *testing.T) {
	logger, err := Initialize("")
	require.NoError(t, err, "Expected Initialize to succeed with empty logging level")
	require.NotNil(t, logger, "Expected Initialize to return a non-nil logger")

	expectedLevel, err := zap.ParseAtomicLevel("info")
	require.NoError(t, err, "Expected ParseAtomicLevel to succeed")
	assert.Equal(t, expectedLevel.Level(), logger.Level(), "Expected logger level to match the specified level")
}
