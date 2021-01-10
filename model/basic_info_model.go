package model

import (
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
	"time"
)

type BasicInfo struct {
	Uid        uint64 `gorm:"primary_key;index:idx"`
	Nickname   string `gorm:"type:varchar(32)"`
	State      uint8
	Expiration uint64
	Email      string    `gorm:"type:varchar(64)"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

var db *gorm.DB

func LoadBasicInfoModel(gormDB *gorm.DB) {
	if !gormDB.HasTable(&BasicInfo{}) {
		logger.Warn("Table basic_info not found, start creating")
		if err := gormDB.Set(
			"gorm:table_options",
			"ENGINE=Innodb DEFAULT CHARSET=utf8mb4",
		).CreateTable(&BasicInfo{}).Error; err != nil {
			logger.Painc("Create table basic_info failed: " + err.Error())
		}
	}
	db = gormDB
}

func (basic *BasicInfo) Get() error {
	return db.Find(&basic, "`email` = ?", basic.Email).Error
}

func (basic *BasicInfo) Save() error {
	return db.Save(&basic).Error
}