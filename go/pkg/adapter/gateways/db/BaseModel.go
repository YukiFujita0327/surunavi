package db

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID			uint 			`gorm:"primaryKey"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeleteAt	gorm.DeletedAt	`gorm:"index"`
}
