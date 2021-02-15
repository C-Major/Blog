package util

import (
	"testing"
)

func TestIsTestEnv(t *testing.T) {
	t.Logf("Test Env: %v", IsTestEnv())
}
