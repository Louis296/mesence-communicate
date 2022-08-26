package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	Id         int       `gorm:"id"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
	Status     int       `gorm:"status"`
}

func (b *Base) BeforeCreate(db *gorm.DB) error {
	b.CreateTime = time.Now()
	b.UpdateTime = time.Now()
	return nil
}

func (b *Base) BeforeUpdate(db *gorm.DB) error {
	b.UpdateTime = time.Now()
	return nil
}
