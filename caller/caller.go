package caller

import (
	"fmt"

	"github.com/c-major/blog/common"
	"github.com/c-major/blog/constdef"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DBRead .
	DBRead *gorm.DB

	// DBWrite .
	DBWrite *gorm.DB
)

// InitCaller .
func InitCaller(config *common.Config) error {
	err := initDBRead(config)
	if err != nil {
		common.TextLog.Error("[InitCaller] failed to initialize read db")
		return err
	}

	err = initDBWrite(config)
	if err != nil {
		common.TextLog.Error("[InitCaller] failed to initialize write db")
		return err
	}

	return nil
}

func initDBRead(config *common.Config) error {
	dsn := fmt.Sprintf(constdef.DSNTemplate,
		config.DBConfig.DBReadConfig.Username,
		config.DBConfig.DBReadConfig.Password,
		config.DBConfig.DBReadConfig.Host,
		config.DBConfig.DBReadConfig.Port,
		config.DBConfig.DBReadConfig.Name,
	)

	var err error
	DBRead, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		common.TextLog.Error("[initDBRead] failed to initialize db session")
		return err
	}

	return nil
}

func initDBWrite(config *common.Config) error {
	dsn := fmt.Sprintf(constdef.DSNTemplate,
		config.DBConfig.DBWriteConfig.Username,
		config.DBConfig.DBWriteConfig.Password,
		config.DBConfig.DBWriteConfig.Host,
		config.DBConfig.DBWriteConfig.Port,
		config.DBConfig.DBWriteConfig.Name,
	)

	var err error
	DBWrite, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		common.TextLog.Error("[initDBWrite] failed to initialize db session")
		return err
	}

	return nil
}
