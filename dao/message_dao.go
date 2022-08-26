package dao

import (
	"github.com/louis296/mesence-communicate/dao/model"
	"gorm.io/gorm"
)

func CreateMessage(message model.Message, tx *gorm.DB) error {
	tx = tx.Create(&message)
	return tx.Error
}
