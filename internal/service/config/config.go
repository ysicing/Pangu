package config

import (
	"gitea.ysicing.net/cloud/pangu/common"
	"gitea.ysicing.net/cloud/pangu/internal/models/config"
	"gorm.io/gorm"
)

// Init 检查是否初始化过
func Init() bool {
	status, err := config.Get(common.InitKey)
	if err != nil {
		return false
	}
	return status == common.InitValue
}

func InitDone() error {
	_, err := config.Get(common.InitKey)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return config.Set(common.InitKey, common.InitValue)
}
