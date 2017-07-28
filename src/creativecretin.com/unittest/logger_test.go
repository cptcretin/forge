package unittest

import (
    "testing"

    "creativecretin.com/forge/1.0/logger"
)

func TestLogWriter(t *testing.T) {
    logger.Trace("This is a trace message")
    logger.Info("This is an info message")
    logger.Warn("This is a warning message")
    logger.Error("This is an error message")
}