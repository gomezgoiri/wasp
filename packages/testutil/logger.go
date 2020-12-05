package testutil

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"testing"

	"github.com/iotaledger/hive.go/logger"
	"go.uber.org/zap"
)

// NewLogger produces a logger adjusted for test cases.
func NewLogger(t *testing.T) *logger.Logger {
	log, err := zap.NewDevelopment()
	require.NoError(t, err)
	return log.Named(t.Name()).Sugar()
}

// WithLevel returns a logger with a level increased.
// Can be useful in tests to disable logging in some parts of the system.
func WithLevel(log *logger.Logger, level logger.Level, printStackTrace bool) *logger.Logger {
	if !printStackTrace {
		return log.Desugar().WithOptions(zap.IncreaseLevel(level), zap.AddStacktrace(zapcore.PanicLevel)).Sugar()
	}
	return log.Desugar().WithOptions(zap.IncreaseLevel(level)).Sugar()
}