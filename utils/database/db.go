package database

import (
	"github.com/jinzhu/gorm"
	"github.com/szdx4/attsys-server/config"
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
