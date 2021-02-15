package common

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/c-major/blog/constdef"
	"github.com/c-major/blog/util"
	"gopkg.in/yaml.v2"
)

// Config .
type Config struct {
	DBConfig *DBConfig `yaml:"database"`
}

// DBConfig .
type DBConfig struct {
	DBReadConfig  *DBReadConfig  `yaml:"read"`
	DBWriteConfig *DBWriteConfig `yaml:"write"`
}

// DBReadConfig .
type DBReadConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// DBWriteConfig .
type DBWriteConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// GetConfig .
func GetConfig(dirs ...string) (*Config, error) {
	if len(dirs) == 0 {
		TextLog.Error("[InitConfig] dirs are empty")
		return nil, errors.New("dirs are empty")
	}

	confName := constdef.AppName
	if util.IsTestEnv() {
		confName += "_test"
	}
	confName += ".yaml"
	confPath := filepath.Join(dirs...) + "/" + confName
	confBytes, err := ioutil.ReadFile(confPath)
	if err != nil {
		TextLog.WithField("confPath", confPath).Errorf("[InitConfig] failed to load conf file, err=%s", err)
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(confBytes, &config)
	if err != nil {
		TextLog.WithField("confBytes", confBytes).Errorf("[InitConfig] failed to unmarshal conf file, err=%s", err)
		return nil, err
	}

	return &config, nil
}
