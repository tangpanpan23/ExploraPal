package hps

import (
	"time"

	"explorapal/app/model"
)

// User 用户表
type User struct {
	model.BaseModel

	UserID      int64  `gorm:"column:user_id;uniqueIndex;not null;comment:用户ID"`
	Username    string `gorm:"column:username;size:50;comment:用户名"`
	Nickname    string `gorm:"column:nickname;size:100;comment:昵称"`
	Avatar      string `gorm:"column:avatar;size:500;comment:头像URL"`
	Age         int32  `gorm:"column:age;comment:年龄"`
	Gender      string `gorm:"column:gender;size:10;comment:性别：male,female,other"`
	Phone       string `gorm:"column:phone;size:20;comment:手机号"`
	Email       string `gorm:"column:email;size:100;comment:邮箱"`
	Status      string `gorm:"column:status;size:20;default:active;comment:状态：active,inactive"`
	LastLoginAt *time.Time `gorm:"column:last_login_at;comment:最后登录时间"`
}

// TableName 设置表名
func (User) TableName() string {
	return "users"
}
