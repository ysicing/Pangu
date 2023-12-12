package config

import (
	"gitea.ysicing.net/cloud/pangu/internal/db"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	Ckey string `gorm:"unique;not null;index" json:"ckey"`
	Cval string `gorm:"unique;" json:"cval"`
}

func (c *Config) TableName() string {
	return "config"
}

func init() {
	db.Migrate(Config{})
}

// FindAll 查询所有
func FindAll() (o []Config) {
	if db.DB.Find(&o).Error != nil {
		return nil
	}
	return
}

// Get 获取配置
func Get(ckey string) (string, error) {
	var obj Config
	has := db.DB.Model(Config{}).Where("ckey=?", ckey).Last(&obj)
	if has.Error != nil && has.Error != gorm.ErrRecordNotFound {
		return "", has.Error
	}

	if has.RowsAffected == 0 {
		return "", nil
	}

	return obj.Cval, nil
}

func Create(o *Config) error {
	return db.DB.Create(o).Error
}

func Save(o *Config) error {
	return db.DB.Where("ckey = ?", o.Ckey).Save(o).Error
}

// Set 添加配置
func Set(ckey, cval string) error {
	var obj Config
	has := db.DB.Model(Config{}).Where("ckey=?", ckey).Last(&obj)
	if has.Error != nil && has.Error != gorm.ErrRecordNotFound {
		return has.Error
	}
	var err error
	if has.RowsAffected == 0 {
		err = Create(&Config{
			Ckey: ckey,
			Cval: cval,
		})
	} else {
		obj.Cval = cval
		err = Save(&obj)
	}
	return err
}
