package model

import "time"

type Message struct {
	ID         int64  `gorm:"primaryKey"`
	FromUserID int64  `gorm:"index;not null"`
	ToUserID   int64  `gorm:"index;not null"`
	Content    string `gorm:"type:text;not null"`
	CreatedAt  time.Time
}
