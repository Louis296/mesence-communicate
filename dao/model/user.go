package model

type User struct {
	Base
	Phone    string `gorm:"column:phone;NOT NULL"`
	Name     string `gorm:"column:name;NOT NULL"`
	Password string `gorm:"column:password;NOT NULL"`
	Avatar   string `gorm:"column:avatar;NOT NULL"`
	Location string `gorm:"column:location;NOT NULL"`
}

func (m *User) TableName() string {
	return "user"
}

type UserResp struct {
	Id       int
	Phone    string
	Name     string
	Avatar   string
	Location string
}

func (m User) GenResp() UserResp {
	return UserResp{
		Id:       m.Id,
		Phone:    m.Phone,
		Name:     m.Name,
		Avatar:   m.Avatar,
		Location: m.Location,
	}
}
