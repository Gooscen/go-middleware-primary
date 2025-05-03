package model

import (
	"go-middleware-primary/config"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

// 提供用户验证函数
func CheckUser(username, password string) bool {
	var user User
	result := config.DB.Where("username = ? AND password = ?", username, password).First(&user)
	return result.Error == nil
}
