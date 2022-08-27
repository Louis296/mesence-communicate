package dao

import (
	"github.com/louis296/mesence-communicate/dao/model"
	"gorm.io/gorm"
)

func CreateFriendRequest(tx *gorm.DB, request model.FriendRequest) error {
	tx = tx.Create(&request)
	return tx.Error
}

func SaveFriendRequest(tx *gorm.DB, request model.FriendRequest) error {
	tx = tx.Save(&request)
	return tx.Error
}

func ListFriendRequestsByCandidate(candidate string, limit, offset int) ([]model.FriendRequest, int64, error) {
	sql := DB
	sql = sql.Model(&model.FriendRequest{})
	var ans []model.FriendRequest
	sql = sql.Where("candidate=?", candidate)
	var total int64
	sql.Count(&total)
	err := sql.Limit(limit).Offset(limit * (offset - 1)).Scan(&ans).Error
	if err != nil {
		return nil, 0, err
	}
	return ans, total, nil
}

func ListFriendRequestsBySender(sender string, limit, offset int) ([]model.FriendRequest, int64, error) {
	sql := DB
	sql = sql.Model(&model.FriendRequest{})
	var ans []model.FriendRequest
	sql = sql.Where("sender=?", sender)
	var total int64
	sql.Count(&total)
	err := sql.Limit(limit).Offset(limit * (offset - 1)).Scan(&ans).Error
	if err != nil {
		return nil, 0, err
	}
	return ans, total, nil
}

func GetFriendRequestById(id int) (*model.FriendRequest, error) {
	sql := DB
	sql = sql.Model(&model.FriendRequest{})
	var ans model.FriendRequest
	err := sql.Where("id=?", id).First(&ans).Error
	if err != nil {
		return nil, err
	}
	return &ans, nil
}
