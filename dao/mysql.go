package dao

import (
	"mini_douyin/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return db.AutoMigrate(&model.User{}, &model.Video{}, &model.Favorite{}, &model.Comment{}, &model.Relation{}, &model.Message{})
}
