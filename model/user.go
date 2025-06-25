package model

import "time"

type User struct {
	ID              int64  `gorm:"primaryKey"`
	Username        string `gorm:"unique;size:32;not null"`
	Password        string `gorm:"size:64;not null"`
	Name            string `gorm:"size:32;not null"`
	Avatar          string
	BackgroundImage string
	Signature       string
	CreatedAt       time.Time
}
