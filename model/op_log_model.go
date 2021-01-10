package model

import (
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
	"time"
)

type OpLog struct {
	Uid       uint64
	ResId     uint64
	Ip        string    `gorm:"type:varchar(32)"`
	OpType    string    `gorm:"type:varchar(32)"`
	Info      string    `gorm:"type:varchar(128)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func LoadOpLogModel(gormDB *gorm.DB) {
	if !gormDB.HasTable(&OpLog{}) {
		logger.Warn("Table op_log not found, start creating")
		if err := gormDB.Set(
			"gorm:table_options",
			"ENGINE=Innodb DEFAULT CHARSET=utf8mb4",
		).CreateTable(&OpLog{}).Error; err != nil {
			logger.Painc("Create table op_log failed: " + err.Error())
		}
	}
}
