package model

import (
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
)

type Resource struct {
	ResId   uint64 `gorm:"primary_key;index:idx"`
	ResKey  string `gorm:"type:varchar(64)"`
	Allow   string `gorm:"type:varchar(128)"`
	Disable string `gorm:"type:varchar(128)"`
	State   uint8
}

func LoadResourceModel(gormDB *gorm.DB) {
	if !gormDB.HasTable(&Resource{}) {
		logger.Warn("Table resource not found, start creating")
		if err := gormDB.Set(
			"gorm:table_options",
			"ENGINE=Innodb DEFAULT CHARSET=utf8mb4",
		).CreateTable(&Resource{}).Error; err != nil {
			logger.Painc("Create table resource failed: " + err.Error())
		}
	}
}
