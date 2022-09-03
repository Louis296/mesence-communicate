package dao

import (
	"github.com/louis296/mesence-communicate/dao/model"
	"github.com/louis296/mesence-communicate/pkg/util"
	"gorm.io/gorm"
)

func CreateMessage(message model.Message, tx *gorm.DB) error {
	tx = tx.Create(&message)
	return tx.Error
}

func ListMessageByTwoUserPhone(userA, userB string, offset, limit int, startTime, endTime string) ([]model.Message, error) {
	sql := DB
	sql = sql.Model(&model.Message{})
	var res []model.Message
	sql = sql.Where("(`from`=? and `to`=?) or (`from`=? and `to`=?)", userA, userB, userB, userA)
	if startTime != "" {
		sql = sql.Where("send_time >= ?", util.TimeParse(startTime))
	}
	if endTime != "" {
		sql = sql.Where("send_time <= ?", util.TimeParse(endTime))
	}
	err := sql.Limit(limit).Offset((offset - 1) * limit).Order("send_time desc").Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
