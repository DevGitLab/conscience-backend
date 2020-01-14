package dao

import (
	"conscience-backend/model"
	"github.com/jinzhu/gorm"
)

func Init(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&model.User{})
}
