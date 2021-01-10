package model

import (
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
	"time"
)

type Account struct {
	Uid          uint64
	Identifier   string `gorm:"type:varchar(32)"`
	Credential   string `gorm:"type:varchar(32)"`
	Name         string `gorm:"type:varchar(32)"`
	Salt         string `gorm:"type:varchar(64)"`
	IdentityType uint8
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

func LoadAccountModel(gormDB *gorm.DB) {
	if !gormDB.HasTable(&Account{}) {
		logger.Warn("Table account not found, start creating")
		if err := gormDB.Set(
			"gorm:table_options",
			"ENGINE=Innodb DEFAULT CHARSET=utf8mb4",
		).CreateTable(&Account{}).Error; err != nil {
			logger.Painc("Create table account failed: " + err.Error())
		}
	}
}

func (account *Account) Save() error {
	return db.Save(&account).Error
}

func (account *Account) ExistByUidAndType(types []uint8) (bool, error) {
	count := uint8(0)
	err := db.Model(&Account{}).Where("`uid` = ? AND `identity_type` IN (?)", account.Uid, types).Count(&count).Error
	return count > 0, err
}
