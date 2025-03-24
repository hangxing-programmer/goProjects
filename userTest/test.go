package userTest

import (
	"abc/dao"
)

type Contact01 struct {
	UserID      int
	ContactID   string
	ContactType string
	Status      string
	UserName    string
	CreateTime  string `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime  string `gorm:"default:CURRENT_TIMESTAMP"`
}

// TableName 结构体对应的数据库表名
func (Contact01) TableName() string {
	return "user_contact"
}

// AddUserById 单个添加
func AddUserById(user *Contact01) {
	dao.Db.Create(user)
}

// UpdateUserById 修改名字
func UpdateUserById(user *Contact01) {
	dao.Db.Model(user).Where("user_id = ?", user.UserID).Update("user_name", user.UserName)
}

// DeleteUserById 删除用户
func DeleteUserById(user Contact01) {
	dao.Db.Where("user_id = ?", user.UserID).Delete(Contact01{})
}

// GetUserInfoTest 查找单个用户
func GetUserInfoTest(id int) interface{} {
	var user Contact01
	dao.Db.Where("user_id = ?", id).First(&user)
	return user
}

// FindUsers 查找多个用户
func FindUsers(user Contact01) []Contact01 {
	var users []Contact01
	dao.Db.Where("user_id < ?", user.UserID).Find(&users)
	return users
}
