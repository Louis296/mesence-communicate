package model

type FriendRelation struct {
	Base
	UserID         int    `gorm:"column:user_id;NOT NULL"`
	UserPhone      string `gorm:"column:user_phone"`
	FriendID       int    `gorm:"column:friend_id;NOT NULL"`
	FriendPhone    string `gorm:"column:friend_phone"`
	FriendNoteName string `gorm:"column:friend_note_name"`
}

func (m *FriendRelation) TableName() string {
	return "friend_relation"
}
