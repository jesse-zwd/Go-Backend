package model

import (
	"gorm.io/gorm"
	"time"
)

type GORM_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"id"` 
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}