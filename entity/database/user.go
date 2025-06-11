package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   int64  // userID
	Username string // 用户名
	Password string // 密码
	Gender   bool   // 性别
	Email    string // 邮箱
}
