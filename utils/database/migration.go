package database

import (
	"github.com/agcx_server/models"
	"golang.org/x/crypto/bcrypt"
)

// Migrate 执行数据库迁移
func Migrate() {
	Connector.AutoMigrate(&models.User{})
	Connector.AutoMigrate(&models.Score{})
}

// Seed 执行数据库填充
func Seed() {
	userCount := 0
	Connector.Model(&models.User{}).Count(&userCount)
	if userCount == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte("root"), 10)
		if err != nil {
			panic(err)
		}

		user := models.User{}
		user.Name = "root"
		user.Email = "root@gmail.com"
		user.Password = string(hash)

		Connector.Save(&user)
	}
}
