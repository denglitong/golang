package testify_demo

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRequire(t *testing.T) {
	// The require package provides same global functions as the assert package,
	// but instead of returning a boolean result they terminate current test.
	// These functions must be called from the goroutine running the test or benchmark function,
	// not from other goroutines created during the test. Otherwise, race conditions may occur.
	require.Equal(t, 123, 123)
}
