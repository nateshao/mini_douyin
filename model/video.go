package model

import "time"

type Video struct {
	ID        int64  `gorm:"primaryKey"`
	AuthorID  int64  `gorm:"index;not null"`
	PlayUrl   string `gorm:"not null"`
	CoverUrl  string `gorm:"not null"`
	Title     string
	CreatedAt time.Time
}
