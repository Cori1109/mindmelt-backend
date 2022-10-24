package models

import (
	"gorm.io/gorm"
)

type News struct {
	NEWS_ID string `gorm:"primary key;serial" json:"news_id"`
	Content *string `json:"content"`
	Createdat string `json:"createdAt"`
}

func MigrateNews(db *gorm.DB) error {
	err := db.AutoMigrate(&News{})

	return err
}