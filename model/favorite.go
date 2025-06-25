package model

import "time"

type Favorite struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64 `gorm:"index;not null"`
	VideoID   int64 `gorm:"index;not null"`
	CreatedAt time.Time
}
