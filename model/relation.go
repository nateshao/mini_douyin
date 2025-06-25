package model

import "time"

type Relation struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64 `gorm:"index;not null"`
	ToUserID  int64 `gorm:"index;not null"`
	CreatedAt time.Time
}
