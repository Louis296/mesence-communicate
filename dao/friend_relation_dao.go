package dao

import "github.com/louis296/mesence-communicate/dao/model"

func GetFriendRelationsByUserPhone(phone string) ([]model.FriendRelation, error) {
	sql := DB
	sql = sql.Model(&model.FriendRelation{})
	var ans []model.FriendRelation
	err := sql.Where("user_phone=?", phone).Scan(&ans).Error
	if err != nil {
		return nil, err
	}
	return ans, nil
}
