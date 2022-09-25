package dao

import "github.com/louis296/mesence-communicate/dao/model"

func GetUserByPhone(phone string) (*model.User, error) {
	sql := DB
	ans := &model.User{}
	sql = sql.Model(&model.User{})
	err := sql.Where("phone=?", phone).First(&ans).Error
	if err != nil {
		return nil, err
	}
	return ans, nil
}

func GetUserIdLookupByIds(userIds []int) (map[int]model.User, error) {
	sql := DB
	var res []model.User
	sql = sql.Model(&model.User{})
	err := sql.Where("id in ?", userIds).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	ans := make(map[int]model.User)
	for _, item := range res {
		ans[item.Id] = item
	}
	return ans, nil
}
