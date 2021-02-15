package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	InitLog()
}

func TestGetConfig(t *testing.T) {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	os.Setenv("IS_TEST_ENV", "1")
	config, err := GetConfig(rootDir, "../", "conf")
	assert.Nil(t, err)
	t.Logf("config.DBConfig.DBReadConfig: %+v", config.DBConfig.DBReadConfig)
	t.Logf("config.DBConfig.DBWriteConfig: %+v", config.DBConfig.DBWriteConfig)
}
