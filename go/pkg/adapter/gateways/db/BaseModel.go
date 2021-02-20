package db

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id			string 			`gorm:"primaryKey;column:ID"`
	CreatedAt	time.Time		`gorm:"column:CREATED_AT"`
	UpdatedAt	time.Time		`gorm:"column:UPDATED_AT"`
	DeletedAt	gorm.DeletedAt	`gorm:"index;column:DELETED_AT"`
}
