package caller

import (
	"os"
	"testing"

	"github.com/c-major/blog/common"
	"github.com/stretchr/testify/assert"
)

func init() {
	common.InitLog()
}

func TestInitCaller(t *testing.T) {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	os.Setenv("IS_TEST_ENV", "1")
	config, err := common.GetConfig(rootDir, "../", "conf")
	err = InitCaller(config)
	assert.Nil(t, err)
}
