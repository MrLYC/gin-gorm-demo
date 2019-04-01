package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mrlyc/gin-gorm-demo/config"
)

// DB :
var DB *gorm.DB

// OpenDB :
func OpenDB() error {
	db, err := gorm.Open(config.Configuration.Database.Type, config.Configuration.Database.Connection)
	if err != nil {
		return err
	}
	DB = db
	return nil
}

// CloseDB :
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
