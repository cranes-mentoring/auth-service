package model

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
}
