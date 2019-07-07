package database

import (
	"github.com/agcx_server/config"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jinzhu/gorm"
	"log"
)

// Connector 数据库连接器
var Connector *gorm.DB

// Connect 连接数据库
func Connect() {
	var err error
	Connector, err = gorm.Open(config.Database.Type, config.Database.ConnectionString())
	if err != nil {
		log.Fatalf("不能连接到数据库: %v", err)
	}
}
