package dao

import "github.com/louis296/mesence-communicate/dao/model"

func GetUserByUserPhone(phone string) (*model.User, error) {
	sql := DB
	ans := &model.User{}
	sql = sql.Model(&model.User{})
	err := sql.Where("phone=?", phone).First(&ans).Error
	if err != nil {
		return nil, err
	}
	return ans, nil
}
