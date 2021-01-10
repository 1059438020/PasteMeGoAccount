package model

import (
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
	"time"
)

type Permission struct {
	Uid       uint64
	ResId     uint64
	Type      uint8
	StartTime time.Time
	EndTime   time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func LoadPermissionModel(gormDB *gorm.DB) {
	if !gormDB.HasTable(&Permission{}) {
		logger.Warn("Table permission not found, start creating")
		if err := gormDB.Set(
			"gorm:table_options",
			"ENGINE=Innodb DEFAULT CHARSET=utf8mb4",
		).CreateTable(&Permission{}).Error; err != nil {
			logger.Painc("Create table permission failed: " + err.Error())
		}
	}
}
