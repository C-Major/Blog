package util

import "os"

// IsTestEnv checks whether it is a test environment
func IsTestEnv() bool {
	return os.Getenv("IS_TEST_ENV") == "1"
}
