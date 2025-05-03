package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/go_middleware_primary?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 自动迁移表
	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("自动建表失败:", err)
	}
}

// 为了避免循环导入，这里引入模型在最后
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}
